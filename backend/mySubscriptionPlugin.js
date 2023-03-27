import { makeExtendSchemaPlugin, gql, embed } from "graphile-utils";

const newSpacePostsTopic = async(args, context, _resolveInfo) => {
  if (!context.personId && !context.userId) {
    throw new Error("No personId or userId in context");
  }
  return `graphql:spaceposts:${args.spaceId}`
};

export default makeExtendSchemaPlugin(({ pgSql: sql }) => ({
  typeDefs: gql`
    type PostSubscriptionPayload {
      post: Post
      event: String
    }

    extend type Subscription {
      posts(spaceId: UUID!): PostSubscriptionPayload @pgSubscription(topic: ${embed(newSpacePostsTopic)})
    }
  `,

  resolvers: {
    PostSubscriptionPayload: {
      // This method finds the post from the database based on the event
      // published by PostgreSQL.
      //
      // In a future release, we hope to enable you to replace this entire
      // method with a small schema directive above, should you so desire. It's
      // mostly boilerplate.
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
