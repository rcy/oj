// MySubscriptionPlugin.js
import { makeExtendSchemaPlugin, gql, embed } from "graphile-utils";
// or: import { makeExtendSchemaPlugin, gql, embed } from 'graphile-utils';

const currentUserTopicFromContext = async (_args, context, _resolveInfo) => {
  //console.log('oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo', context)
  if (context.userId) {
    const topic = `graphql:user:${context.userId}`;
    console.log({ topic })
    return topic
  } else {
    throw new Error("You're not logged in");
  }
};

const newSpacePostsTopic = async(args, context, _resolveInfo) => {
  console.log('xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx', { args, context })
  return `graphql:spaceposts:${args.spaceId}`
};

export default makeExtendSchemaPlugin(({ pgSql: sql }) => ({
  typeDefs: gql`
    type UserSubscriptionPayload {
      # This is populated by our resolver below
      user: User

      # This is returned directly from the PostgreSQL subscription payload (JSON object)
      event: String
    }

    type PostSubscriptionPayload {
      post: Post
      event: String
    }

    extend type Subscription {
      """
      Triggered when the current user's data changes:

      - direct modifications to the user
      - when their organization membership changes
      """
      currentUserUpdated: UserSubscriptionPayload @pgSubscription(topic: ${embed(currentUserTopicFromContext)})
      
      posts(spaceId: UUID!): PostSubscriptionPayload @pgSubscription(topic: ${embed(newSpacePostsTopic)})
    }
  `,

  resolvers: {
    UserSubscriptionPayload: {
      // This method finds the user from the database based on the event
      // published by PostgreSQL.
      //
      // In a future release, we hope to enable you to replace this entire
      // method with a small schema directive above, should you so desire. It's
      // mostly boilerplate.
      async user(
        event,
        _args,
        _context,
        { graphile: { selectGraphQLResultFromTable } }
      ) {
        const rows = await selectGraphQLResultFromTable(
          sql.fragment`app_public.users`,
          (tableAlias, sqlBuilder) => {
            sqlBuilder.where(
              sql.fragment`${tableAlias}.id = ${sql.value(event.subject)}`
            );
          }
        );
        return rows[0];
      },
    },
    PostSubscriptionPayload: {
      async post(
        event,
        _args,
        _context,
        { graphile: { selectGraphQLResultFromTable } }
      ) {
        const rows = await selectGraphQLResultFromTable(
          sql.fragment`app_public.posts`,
          (tableAlias, sqlBuilder) => {
            sqlBuilder.where(
              sql.fragment`${tableAlias}.id = ${sql.value(event.subject)}`
            );
          }
        );
        return rows[0];
      },
    },
  },
}));
