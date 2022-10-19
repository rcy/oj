import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  /** A location in a connection that can be used for resuming pagination. */
  Cursor: any;
  /**
   * A point in time as described by the [ISO
   * 8601](https://en.wikipedia.org/wiki/ISO_8601) standard. May or may not include a timezone.
   */
  Datetime: any;
  /** The `JSON` scalar type represents JSON values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf). */
  JSON: any;
  /** A universally unique identifier as defined by [RFC 4122](https://tools.ietf.org/html/rfc4122). */
  UUID: any;
};

export type Authentication = Node & {
  __typename?: 'Authentication';
  createdAt: Scalars['Datetime'];
  details: Scalars['JSON'];
  id: Scalars['UUID'];
  identifier: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  service: Scalars['String'];
  updatedAt: Scalars['Datetime'];
  /** Reads a single `User` that is related to this `Authentication`. */
  user?: Maybe<User>;
  userId: Scalars['UUID'];
};

/**
 * A condition to be used against `Authentication` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type AuthenticationCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `details` field. */
  details?: InputMaybe<Scalars['JSON']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `identifier` field. */
  identifier?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `service` field. */
  service?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `userId` field. */
  userId?: InputMaybe<Scalars['UUID']>;
};

/** A connection to a list of `Authentication` values. */
export type AuthenticationsConnection = {
  __typename?: 'AuthenticationsConnection';
  /** A list of edges which contains the `Authentication` and cursor to aid in pagination. */
  edges: Array<AuthenticationsEdge>;
  /** A list of `Authentication` objects. */
  nodes: Array<Authentication>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Authentication` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Authentication` edge in the connection. */
export type AuthenticationsEdge = {
  __typename?: 'AuthenticationsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Authentication` at the end of the edge. */
  node: Authentication;
};

/** Methods to use when ordering `Authentication`. */
export enum AuthenticationsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  DetailsAsc = 'DETAILS_ASC',
  DetailsDesc = 'DETAILS_DESC',
  IdentifierAsc = 'IDENTIFIER_ASC',
  IdentifierDesc = 'IDENTIFIER_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  ServiceAsc = 'SERVICE_ASC',
  ServiceDesc = 'SERVICE_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC',
  UserIdAsc = 'USER_ID_ASC',
  UserIdDesc = 'USER_ID_DESC'
}

/** All input for the create `Family` mutation. */
export type CreateFamilyInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Family` to be created by this mutation. */
  family: FamilyInput;
};

/** All input for the create `FamilyMembership` mutation. */
export type CreateFamilyMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `FamilyMembership` to be created by this mutation. */
  familyMembership: FamilyMembershipInput;
};

/** The output of our create `FamilyMembership` mutation. */
export type CreateFamilyMembershipPayload = {
  __typename?: 'CreateFamilyMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  /** The `FamilyMembership` that was created by this mutation. */
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our create `FamilyMembership` mutation. */
export type CreateFamilyMembershipPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** The output of our create `Family` mutation. */
export type CreateFamilyPayload = {
  __typename?: 'CreateFamilyPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Family` that was created by this mutation. */
  family?: Maybe<Family>;
  /** An edge for our `Family`. May be used by Relay 1. */
  familyEdge?: Maybe<FamiliesEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `User` that is related to this `Family`. */
  user?: Maybe<User>;
};


/** The output of our create `Family` mutation. */
export type CreateFamilyPayloadFamilyEdgeArgs = {
  orderBy?: InputMaybe<Array<FamiliesOrderBy>>;
};

/** All input for the create `Interest` mutation. */
export type CreateInterestInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Interest` to be created by this mutation. */
  interest: InterestInput;
};

/** The output of our create `Interest` mutation. */
export type CreateInterestPayload = {
  __typename?: 'CreateInterestPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Interest` that was created by this mutation. */
  interest?: Maybe<Interest>;
  /** An edge for our `Interest`. May be used by Relay 1. */
  interestEdge?: Maybe<InterestsEdge>;
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
};


/** The output of our create `Interest` mutation. */
export type CreateInterestPayloadInterestEdgeArgs = {
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** All input for the `createNewFamilyMember` mutation. */
export type CreateNewFamilyMemberInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  name: Scalars['String'];
  role: Scalars['String'];
};

/** The output of our `createNewFamilyMember` mutation. */
export type CreateNewFamilyMemberPayload = {
  __typename?: 'CreateNewFamilyMemberPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our `createNewFamilyMember` mutation. */
export type CreateNewFamilyMemberPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the create `Person` mutation. */
export type CreatePersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Person` to be created by this mutation. */
  person: PersonInput;
};

/** The output of our create `Person` mutation. */
export type CreatePersonPayload = {
  __typename?: 'CreatePersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Person` that was created by this mutation. */
  person?: Maybe<Person>;
  /** An edge for our `Person`. May be used by Relay 1. */
  personEdge?: Maybe<PeopleEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our create `Person` mutation. */
export type CreatePersonPayloadPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};

/** All input for the create `Post` mutation. */
export type CreatePostInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Post` to be created by this mutation. */
  post: PostInput;
};

/** The output of our create `Post` mutation. */
export type CreatePostPayload = {
  __typename?: 'CreatePostPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `SpaceMembership` that is related to this `Post`. */
  membership?: Maybe<SpaceMembership>;
  /** The `Post` that was created by this mutation. */
  post?: Maybe<Post>;
  /** An edge for our `Post`. May be used by Relay 1. */
  postEdge?: Maybe<PostsEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our create `Post` mutation. */
export type CreatePostPayloadPostEdgeArgs = {
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};

/** All input for the create `Space` mutation. */
export type CreateSpaceInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Space` to be created by this mutation. */
  space: SpaceInput;
};

/** All input for the create `SpaceMembership` mutation. */
export type CreateSpaceMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `SpaceMembership` to be created by this mutation. */
  spaceMembership: SpaceMembershipInput;
};

/** The output of our create `SpaceMembership` mutation. */
export type CreateSpaceMembershipPayload = {
  __typename?: 'CreateSpaceMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  /** The `SpaceMembership` that was created by this mutation. */
  spaceMembership?: Maybe<SpaceMembership>;
  /** An edge for our `SpaceMembership`. May be used by Relay 1. */
  spaceMembershipEdge?: Maybe<SpaceMembershipsEdge>;
};


/** The output of our create `SpaceMembership` mutation. */
export type CreateSpaceMembershipPayloadSpaceMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** The output of our create `Space` mutation. */
export type CreateSpacePayload = {
  __typename?: 'CreateSpacePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Space` that was created by this mutation. */
  space?: Maybe<Space>;
  /** An edge for our `Space`. May be used by Relay 1. */
  spaceEdge?: Maybe<SpacesEdge>;
};


/** The output of our create `Space` mutation. */
export type CreateSpacePayloadSpaceEdgeArgs = {
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};

/** All input for the create `Topic` mutation. */
export type CreateTopicInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Topic` to be created by this mutation. */
  topic: TopicInput;
};

/** The output of our create `Topic` mutation. */
export type CreateTopicPayload = {
  __typename?: 'CreateTopicPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Topic` that was created by this mutation. */
  topic?: Maybe<Topic>;
  /** An edge for our `Topic`. May be used by Relay 1. */
  topicEdge?: Maybe<TopicsEdge>;
};


/** The output of our create `Topic` mutation. */
export type CreateTopicPayloadTopicEdgeArgs = {
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};

/** All input for the `deleteFamilyMembershipByNodeId` mutation. */
export type DeleteFamilyMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `FamilyMembership` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteFamilyMembership` mutation. */
export type DeleteFamilyMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `FamilyMembership` mutation. */
export type DeleteFamilyMembershipPayload = {
  __typename?: 'DeleteFamilyMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedFamilyMembershipNodeId?: Maybe<Scalars['ID']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  /** The `FamilyMembership` that was deleted by this mutation. */
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our delete `FamilyMembership` mutation. */
export type DeleteFamilyMembershipPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the `deleteInterestByNodeId` mutation. */
export type DeleteInterestByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Interest` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteInterest` mutation. */
export type DeleteInterestInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Interest` mutation. */
export type DeleteInterestPayload = {
  __typename?: 'DeleteInterestPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedInterestNodeId?: Maybe<Scalars['ID']>;
  /** The `Interest` that was deleted by this mutation. */
  interest?: Maybe<Interest>;
  /** An edge for our `Interest`. May be used by Relay 1. */
  interestEdge?: Maybe<InterestsEdge>;
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
};


/** The output of our delete `Interest` mutation. */
export type DeleteInterestPayloadInterestEdgeArgs = {
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** All input for the `deletePersonByNodeId` mutation. */
export type DeletePersonByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Person` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deletePerson` mutation. */
export type DeletePersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Person` mutation. */
export type DeletePersonPayload = {
  __typename?: 'DeletePersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedPersonNodeId?: Maybe<Scalars['ID']>;
  /** The `Person` that was deleted by this mutation. */
  person?: Maybe<Person>;
  /** An edge for our `Person`. May be used by Relay 1. */
  personEdge?: Maybe<PeopleEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our delete `Person` mutation. */
export type DeletePersonPayloadPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};

/** All input for the `deletePostByNodeId` mutation. */
export type DeletePostByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Post` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deletePost` mutation. */
export type DeletePostInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Post` mutation. */
export type DeletePostPayload = {
  __typename?: 'DeletePostPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedPostNodeId?: Maybe<Scalars['ID']>;
  /** Reads a single `SpaceMembership` that is related to this `Post`. */
  membership?: Maybe<SpaceMembership>;
  /** The `Post` that was deleted by this mutation. */
  post?: Maybe<Post>;
  /** An edge for our `Post`. May be used by Relay 1. */
  postEdge?: Maybe<PostsEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our delete `Post` mutation. */
export type DeletePostPayloadPostEdgeArgs = {
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};

/** All input for the `deleteSpaceByNodeId` mutation. */
export type DeleteSpaceByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Space` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteSpace` mutation. */
export type DeleteSpaceInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** All input for the `deleteSpaceMembershipByNodeId` mutation. */
export type DeleteSpaceMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `SpaceMembership` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteSpaceMembership` mutation. */
export type DeleteSpaceMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `SpaceMembership` mutation. */
export type DeleteSpaceMembershipPayload = {
  __typename?: 'DeleteSpaceMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedSpaceMembershipNodeId?: Maybe<Scalars['ID']>;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  /** The `SpaceMembership` that was deleted by this mutation. */
  spaceMembership?: Maybe<SpaceMembership>;
  /** An edge for our `SpaceMembership`. May be used by Relay 1. */
  spaceMembershipEdge?: Maybe<SpaceMembershipsEdge>;
};


/** The output of our delete `SpaceMembership` mutation. */
export type DeleteSpaceMembershipPayloadSpaceMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** The output of our delete `Space` mutation. */
export type DeleteSpacePayload = {
  __typename?: 'DeleteSpacePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedSpaceNodeId?: Maybe<Scalars['ID']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Space` that was deleted by this mutation. */
  space?: Maybe<Space>;
  /** An edge for our `Space`. May be used by Relay 1. */
  spaceEdge?: Maybe<SpacesEdge>;
};


/** The output of our delete `Space` mutation. */
export type DeleteSpacePayloadSpaceEdgeArgs = {
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};

/** All input for the `deleteTopicByNodeId` mutation. */
export type DeleteTopicByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Topic` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteTopic` mutation. */
export type DeleteTopicInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Topic` mutation. */
export type DeleteTopicPayload = {
  __typename?: 'DeleteTopicPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedTopicNodeId?: Maybe<Scalars['ID']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Topic` that was deleted by this mutation. */
  topic?: Maybe<Topic>;
  /** An edge for our `Topic`. May be used by Relay 1. */
  topicEdge?: Maybe<TopicsEdge>;
};


/** The output of our delete `Topic` mutation. */
export type DeleteTopicPayloadTopicEdgeArgs = {
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};

/** A connection to a list of `Family` values. */
export type FamiliesConnection = {
  __typename?: 'FamiliesConnection';
  /** A list of edges which contains the `Family` and cursor to aid in pagination. */
  edges: Array<FamiliesEdge>;
  /** A list of `Family` objects. */
  nodes: Array<Family>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Family` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Family` edge in the connection. */
export type FamiliesEdge = {
  __typename?: 'FamiliesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Family` at the end of the edge. */
  node: Family;
};

/** Methods to use when ordering `Family`. */
export enum FamiliesOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC',
  UserIdAsc = 'USER_ID_ASC',
  UserIdDesc = 'USER_ID_DESC'
}

export type Family = Node & {
  __typename?: 'Family';
  createdAt: Scalars['Datetime'];
  /** Reads and enables pagination through a set of `FamilyMembership`. */
  familyMemberships: FamilyMembershipsConnection;
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  updatedAt: Scalars['Datetime'];
  /** Reads a single `User` that is related to this `Family`. */
  user?: Maybe<User>;
  userId: Scalars['UUID'];
};


export type FamilyFamilyMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyMembershipCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** A condition to be used against `Family` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type FamilyCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `userId` field. */
  userId?: InputMaybe<Scalars['UUID']>;
};

/** An input for mutations affecting `Family` */
export type FamilyInput = {
  userId: Scalars['UUID'];
};

export type FamilyMembership = Node & {
  __typename?: 'FamilyMembership';
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  familyId: Scalars['UUID'];
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  personId: Scalars['UUID'];
  role: Scalars['String'];
  title?: Maybe<Scalars['String']>;
};

/**
 * A condition to be used against `FamilyMembership` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type FamilyMembershipCondition = {
  /** Checks for equality with the object’s `familyId` field. */
  familyId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `role` field. */
  role?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `title` field. */
  title?: InputMaybe<Scalars['String']>;
};

/** An input for mutations affecting `FamilyMembership` */
export type FamilyMembershipInput = {
  familyId: Scalars['UUID'];
  id?: InputMaybe<Scalars['UUID']>;
  personId: Scalars['UUID'];
  role: Scalars['String'];
  title?: InputMaybe<Scalars['String']>;
};

/** Represents an update to a `FamilyMembership`. Fields that are set will be updated. */
export type FamilyMembershipPatch = {
  familyId?: InputMaybe<Scalars['UUID']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  role?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
};

/** A connection to a list of `FamilyMembership` values. */
export type FamilyMembershipsConnection = {
  __typename?: 'FamilyMembershipsConnection';
  /** A list of edges which contains the `FamilyMembership` and cursor to aid in pagination. */
  edges: Array<FamilyMembershipsEdge>;
  /** A list of `FamilyMembership` objects. */
  nodes: Array<FamilyMembership>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `FamilyMembership` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `FamilyMembership` edge in the connection. */
export type FamilyMembershipsEdge = {
  __typename?: 'FamilyMembershipsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `FamilyMembership` at the end of the edge. */
  node: FamilyMembership;
};

/** Methods to use when ordering `FamilyMembership`. */
export enum FamilyMembershipsOrderBy {
  FamilyIdAsc = 'FAMILY_ID_ASC',
  FamilyIdDesc = 'FAMILY_ID_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  RoleAsc = 'ROLE_ASC',
  RoleDesc = 'ROLE_DESC',
  TitleAsc = 'TITLE_ASC',
  TitleDesc = 'TITLE_DESC'
}

export type FamilyRole = Node & {
  __typename?: 'FamilyRole';
  id: Scalars['Int'];
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
};

/**
 * A condition to be used against `FamilyRole` object types. All fields are tested
 * for equality and combined with a logical ‘and.’
 */
export type FamilyRoleCondition = {
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['Int']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
};

/** A connection to a list of `FamilyRole` values. */
export type FamilyRolesConnection = {
  __typename?: 'FamilyRolesConnection';
  /** A list of edges which contains the `FamilyRole` and cursor to aid in pagination. */
  edges: Array<FamilyRolesEdge>;
  /** A list of `FamilyRole` objects. */
  nodes: Array<FamilyRole>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `FamilyRole` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `FamilyRole` edge in the connection. */
export type FamilyRolesEdge = {
  __typename?: 'FamilyRolesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `FamilyRole` at the end of the edge. */
  node: FamilyRole;
};

/** Methods to use when ordering `FamilyRole`. */
export enum FamilyRolesOrderBy {
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC'
}

export type Interest = Node & {
  __typename?: 'Interest';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  personId?: Maybe<Scalars['UUID']>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
  topicId?: Maybe<Scalars['UUID']>;
  updatedAt: Scalars['Datetime'];
};

/**
 * A condition to be used against `Interest` object types. All fields are tested
 * for equality and combined with a logical ‘and.’
 */
export type InterestCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `topicId` field. */
  topicId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** An input for mutations affecting `Interest` */
export type InterestInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  topicId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Interest`. Fields that are set will be updated. */
export type InterestPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  topicId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `Interest` values. */
export type InterestsConnection = {
  __typename?: 'InterestsConnection';
  /** A list of edges which contains the `Interest` and cursor to aid in pagination. */
  edges: Array<InterestsEdge>;
  /** A list of `Interest` objects. */
  nodes: Array<Interest>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Interest` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Interest` edge in the connection. */
export type InterestsEdge = {
  __typename?: 'InterestsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Interest` at the end of the edge. */
  node: Interest;
};

/** Methods to use when ordering `Interest`. */
export enum InterestsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  TopicIdAsc = 'TOPIC_ID_ASC',
  TopicIdDesc = 'TOPIC_ID_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** The root mutation type which contains root level fields which mutate data. */
export type Mutation = {
  __typename?: 'Mutation';
  /** Creates a single `Family`. */
  createFamily?: Maybe<CreateFamilyPayload>;
  /** Creates a single `FamilyMembership`. */
  createFamilyMembership?: Maybe<CreateFamilyMembershipPayload>;
  /** Creates a single `Interest`. */
  createInterest?: Maybe<CreateInterestPayload>;
  createNewFamilyMember?: Maybe<CreateNewFamilyMemberPayload>;
  /** Creates a single `Person`. */
  createPerson?: Maybe<CreatePersonPayload>;
  /** Creates a single `Post`. */
  createPost?: Maybe<CreatePostPayload>;
  /** Creates a single `Space`. */
  createSpace?: Maybe<CreateSpacePayload>;
  /** Creates a single `SpaceMembership`. */
  createSpaceMembership?: Maybe<CreateSpaceMembershipPayload>;
  /** Creates a single `Topic`. */
  createTopic?: Maybe<CreateTopicPayload>;
  /** Deletes a single `FamilyMembership` using a unique key. */
  deleteFamilyMembership?: Maybe<DeleteFamilyMembershipPayload>;
  /** Deletes a single `FamilyMembership` using its globally unique id. */
  deleteFamilyMembershipByNodeId?: Maybe<DeleteFamilyMembershipPayload>;
  /** Deletes a single `Interest` using a unique key. */
  deleteInterest?: Maybe<DeleteInterestPayload>;
  /** Deletes a single `Interest` using its globally unique id. */
  deleteInterestByNodeId?: Maybe<DeleteInterestPayload>;
  /** Deletes a single `Person` using a unique key. */
  deletePerson?: Maybe<DeletePersonPayload>;
  /** Deletes a single `Person` using its globally unique id. */
  deletePersonByNodeId?: Maybe<DeletePersonPayload>;
  /** Deletes a single `Post` using a unique key. */
  deletePost?: Maybe<DeletePostPayload>;
  /** Deletes a single `Post` using its globally unique id. */
  deletePostByNodeId?: Maybe<DeletePostPayload>;
  /** Deletes a single `Space` using a unique key. */
  deleteSpace?: Maybe<DeleteSpacePayload>;
  /** Deletes a single `Space` using its globally unique id. */
  deleteSpaceByNodeId?: Maybe<DeleteSpacePayload>;
  /** Deletes a single `SpaceMembership` using a unique key. */
  deleteSpaceMembership?: Maybe<DeleteSpaceMembershipPayload>;
  /** Deletes a single `SpaceMembership` using its globally unique id. */
  deleteSpaceMembershipByNodeId?: Maybe<DeleteSpaceMembershipPayload>;
  /** Deletes a single `Topic` using a unique key. */
  deleteTopic?: Maybe<DeleteTopicPayload>;
  /** Deletes a single `Topic` using its globally unique id. */
  deleteTopicByNodeId?: Maybe<DeleteTopicPayload>;
  /** Updates a single `FamilyMembership` using a unique key and a patch. */
  updateFamilyMembership?: Maybe<UpdateFamilyMembershipPayload>;
  /** Updates a single `FamilyMembership` using its globally unique id and a patch. */
  updateFamilyMembershipByNodeId?: Maybe<UpdateFamilyMembershipPayload>;
  /** Updates a single `Interest` using a unique key and a patch. */
  updateInterest?: Maybe<UpdateInterestPayload>;
  /** Updates a single `Interest` using its globally unique id and a patch. */
  updateInterestByNodeId?: Maybe<UpdateInterestPayload>;
  /** Updates a single `Person` using a unique key and a patch. */
  updatePerson?: Maybe<UpdatePersonPayload>;
  /** Updates a single `Person` using its globally unique id and a patch. */
  updatePersonByNodeId?: Maybe<UpdatePersonPayload>;
  /** Updates a single `Post` using a unique key and a patch. */
  updatePost?: Maybe<UpdatePostPayload>;
  /** Updates a single `Post` using its globally unique id and a patch. */
  updatePostByNodeId?: Maybe<UpdatePostPayload>;
  /** Updates a single `Space` using a unique key and a patch. */
  updateSpace?: Maybe<UpdateSpacePayload>;
  /** Updates a single `Space` using its globally unique id and a patch. */
  updateSpaceByNodeId?: Maybe<UpdateSpacePayload>;
  /** Updates a single `SpaceMembership` using a unique key and a patch. */
  updateSpaceMembership?: Maybe<UpdateSpaceMembershipPayload>;
  /** Updates a single `SpaceMembership` using its globally unique id and a patch. */
  updateSpaceMembershipByNodeId?: Maybe<UpdateSpaceMembershipPayload>;
  /** Updates a single `Topic` using a unique key and a patch. */
  updateTopic?: Maybe<UpdateTopicPayload>;
  /** Updates a single `Topic` using its globally unique id and a patch. */
  updateTopicByNodeId?: Maybe<UpdateTopicPayload>;
  /** Updates a single `User` using a unique key and a patch. */
  updateUser?: Maybe<UpdateUserPayload>;
  /** Updates a single `User` using its globally unique id and a patch. */
  updateUserByNodeId?: Maybe<UpdateUserPayload>;
  /** Updates a single `User` using a unique key and a patch. */
  updateUserByPersonId?: Maybe<UpdateUserPayload>;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateFamilyArgs = {
  input: CreateFamilyInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateFamilyMembershipArgs = {
  input: CreateFamilyMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateInterestArgs = {
  input: CreateInterestInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateNewFamilyMemberArgs = {
  input: CreateNewFamilyMemberInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreatePersonArgs = {
  input: CreatePersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreatePostArgs = {
  input: CreatePostInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateSpaceArgs = {
  input: CreateSpaceInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateSpaceMembershipArgs = {
  input: CreateSpaceMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateTopicArgs = {
  input: CreateTopicInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteFamilyMembershipArgs = {
  input: DeleteFamilyMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteFamilyMembershipByNodeIdArgs = {
  input: DeleteFamilyMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteInterestArgs = {
  input: DeleteInterestInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteInterestByNodeIdArgs = {
  input: DeleteInterestByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeletePersonArgs = {
  input: DeletePersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeletePersonByNodeIdArgs = {
  input: DeletePersonByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeletePostArgs = {
  input: DeletePostInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeletePostByNodeIdArgs = {
  input: DeletePostByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceArgs = {
  input: DeleteSpaceInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceByNodeIdArgs = {
  input: DeleteSpaceByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceMembershipArgs = {
  input: DeleteSpaceMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceMembershipByNodeIdArgs = {
  input: DeleteSpaceMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteTopicArgs = {
  input: DeleteTopicInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteTopicByNodeIdArgs = {
  input: DeleteTopicByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateFamilyMembershipArgs = {
  input: UpdateFamilyMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateFamilyMembershipByNodeIdArgs = {
  input: UpdateFamilyMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateInterestArgs = {
  input: UpdateInterestInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateInterestByNodeIdArgs = {
  input: UpdateInterestByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdatePersonArgs = {
  input: UpdatePersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdatePersonByNodeIdArgs = {
  input: UpdatePersonByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdatePostArgs = {
  input: UpdatePostInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdatePostByNodeIdArgs = {
  input: UpdatePostByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceArgs = {
  input: UpdateSpaceInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceByNodeIdArgs = {
  input: UpdateSpaceByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceMembershipArgs = {
  input: UpdateSpaceMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceMembershipByNodeIdArgs = {
  input: UpdateSpaceMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateTopicArgs = {
  input: UpdateTopicInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateTopicByNodeIdArgs = {
  input: UpdateTopicByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateUserByNodeIdArgs = {
  input: UpdateUserByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateUserByPersonIdArgs = {
  input: UpdateUserByPersonIdInput;
};

/** An object with a globally unique `ID`. */
export type Node = {
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
};

/** Information about pagination in a connection. */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** When paginating forwards, the cursor to continue. */
  endCursor?: Maybe<Scalars['Cursor']>;
  /** When paginating forwards, are there more items? */
  hasNextPage: Scalars['Boolean'];
  /** When paginating backwards, are there more items? */
  hasPreviousPage: Scalars['Boolean'];
  /** When paginating backwards, the cursor to continue. */
  startCursor?: Maybe<Scalars['Cursor']>;
};

/** A connection to a list of `Person` values. */
export type PeopleConnection = {
  __typename?: 'PeopleConnection';
  /** A list of edges which contains the `Person` and cursor to aid in pagination. */
  edges: Array<PeopleEdge>;
  /** A list of `Person` objects. */
  nodes: Array<Person>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Person` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Person` edge in the connection. */
export type PeopleEdge = {
  __typename?: 'PeopleEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Person` at the end of the edge. */
  node: Person;
};

/** Methods to use when ordering `Person`. */
export enum PeopleOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

export type Person = Node & {
  __typename?: 'Person';
  createdAt: Scalars['Datetime'];
  /** Reads and enables pagination through a set of `FamilyMembership`. */
  familyMemberships: FamilyMembershipsConnection;
  id: Scalars['UUID'];
  /** Reads and enables pagination through a set of `Interest`. */
  interests: InterestsConnection;
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `SpaceMembership`. */
  spaceMemberships: SpaceMembershipsConnection;
  updatedAt: Scalars['Datetime'];
  /** Reads a single `User` that is related to this `Person`. */
  user?: Maybe<User>;
};


export type PersonFamilyMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyMembershipCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};


export type PersonInterestsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<InterestCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};


export type PersonSpaceMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceMembershipCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** A condition to be used against `Person` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type PersonCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** An input for mutations affecting `Person` */
export type PersonInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name: Scalars['String'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Person`. Fields that are set will be updated. */
export type PersonPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

export type Post = Node & {
  __typename?: 'Post';
  body: Scalars['String'];
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** Reads a single `SpaceMembership` that is related to this `Post`. */
  membership?: Maybe<SpaceMembership>;
  membershipId: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  updatedAt: Scalars['Datetime'];
};

/** A condition to be used against `Post` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type PostCondition = {
  /** Checks for equality with the object’s `body` field. */
  body?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `membershipId` field. */
  membershipId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** An input for mutations affecting `Post` */
export type PostInput = {
  body: Scalars['String'];
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  membershipId: Scalars['UUID'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Post`. Fields that are set will be updated. */
export type PostPatch = {
  body?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  membershipId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `Post` values. */
export type PostsConnection = {
  __typename?: 'PostsConnection';
  /** A list of edges which contains the `Post` and cursor to aid in pagination. */
  edges: Array<PostsEdge>;
  /** A list of `Post` objects. */
  nodes: Array<Post>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Post` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Post` edge in the connection. */
export type PostsEdge = {
  __typename?: 'PostsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Post` at the end of the edge. */
  node: Post;
};

/** Methods to use when ordering `Post`. */
export enum PostsOrderBy {
  BodyAsc = 'BODY_ASC',
  BodyDesc = 'BODY_DESC',
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  MembershipIdAsc = 'MEMBERSHIP_ID_ASC',
  MembershipIdDesc = 'MEMBERSHIP_ID_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** The root query type which gives access points into the data universe. */
export type Query = Node & {
  __typename?: 'Query';
  authentication?: Maybe<Authentication>;
  /** Reads a single `Authentication` using its globally unique `ID`. */
  authenticationByNodeId?: Maybe<Authentication>;
  authenticationByServiceAndIdentifier?: Maybe<Authentication>;
  /** Reads and enables pagination through a set of `Authentication`. */
  authentications?: Maybe<AuthenticationsConnection>;
  currentFamilyMembership?: Maybe<FamilyMembership>;
  currentFamilyMembershipId?: Maybe<Scalars['UUID']>;
  currentUser?: Maybe<User>;
  /** Reads and enables pagination through a set of `Family`. */
  families?: Maybe<FamiliesConnection>;
  family?: Maybe<Family>;
  /** Reads a single `Family` using its globally unique `ID`. */
  familyByNodeId?: Maybe<Family>;
  familyByUserId?: Maybe<Family>;
  familyMembership?: Maybe<FamilyMembership>;
  /** Reads a single `FamilyMembership` using its globally unique `ID`. */
  familyMembershipByNodeId?: Maybe<FamilyMembership>;
  /** Reads and enables pagination through a set of `FamilyMembership`. */
  familyMemberships?: Maybe<FamilyMembershipsConnection>;
  familyRole?: Maybe<FamilyRole>;
  familyRoleByName?: Maybe<FamilyRole>;
  /** Reads a single `FamilyRole` using its globally unique `ID`. */
  familyRoleByNodeId?: Maybe<FamilyRole>;
  /** Reads and enables pagination through a set of `FamilyRole`. */
  familyRoles?: Maybe<FamilyRolesConnection>;
  interest?: Maybe<Interest>;
  /** Reads a single `Interest` using its globally unique `ID`. */
  interestByNodeId?: Maybe<Interest>;
  /** Reads and enables pagination through a set of `Interest`. */
  interests?: Maybe<InterestsConnection>;
  /** Fetches an object given its globally unique `ID`. */
  node?: Maybe<Node>;
  /** The root query type must be a `Node` to work well with Relay 1 mutations. This just resolves to `query`. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `Person`. */
  people?: Maybe<PeopleConnection>;
  person?: Maybe<Person>;
  /** Reads a single `Person` using its globally unique `ID`. */
  personByNodeId?: Maybe<Person>;
  post?: Maybe<Post>;
  /** Reads a single `Post` using its globally unique `ID`. */
  postByNodeId?: Maybe<Post>;
  /** Reads and enables pagination through a set of `Post`. */
  posts?: Maybe<PostsConnection>;
  /**
   * Exposes the root query type nested one level down. This is helpful for Relay 1
   * which can only query top level fields if they are in a particular form.
   */
  query: Query;
  space?: Maybe<Space>;
  /** Reads a single `Space` using its globally unique `ID`. */
  spaceByNodeId?: Maybe<Space>;
  spaceMembership?: Maybe<SpaceMembership>;
  /** Reads a single `SpaceMembership` using its globally unique `ID`. */
  spaceMembershipByNodeId?: Maybe<SpaceMembership>;
  /** Reads and enables pagination through a set of `SpaceMembership`. */
  spaceMemberships?: Maybe<SpaceMembershipsConnection>;
  /** Reads and enables pagination through a set of `Space`. */
  spaces?: Maybe<SpacesConnection>;
  topic?: Maybe<Topic>;
  /** Reads a single `Topic` using its globally unique `ID`. */
  topicByNodeId?: Maybe<Topic>;
  /** Reads and enables pagination through a set of `Topic`. */
  topics?: Maybe<TopicsConnection>;
  user?: Maybe<User>;
  /** Reads a single `User` using its globally unique `ID`. */
  userByNodeId?: Maybe<User>;
  userByPersonId?: Maybe<User>;
  userId?: Maybe<Scalars['UUID']>;
  /** Reads and enables pagination through a set of `User`. */
  users?: Maybe<UsersConnection>;
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationByServiceAndIdentifierArgs = {
  identifier: Scalars['String'];
  service: Scalars['String'];
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<AuthenticationCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<AuthenticationsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryFamiliesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamiliesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyByUserIdArgs = {
  userId: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyMembershipCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRoleArgs = {
  id: Scalars['Int'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRoleByNameArgs = {
  name: Scalars['String'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRoleByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRolesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyRoleCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyRolesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryInterestArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryInterestByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryInterestsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<InterestCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryNodeArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPeopleArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PersonCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryPersonArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPersonByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPostArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPostByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPostsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PostCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceMembershipCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QuerySpacesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryTopicArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryTopicByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryTopicsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<TopicCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryUserArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryUserByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryUserByPersonIdArgs = {
  personId: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<UserCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<UsersOrderBy>>;
};

export type Space = Node & {
  __typename?: 'Space';
  createdAt: Scalars['Datetime'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['UUID'];
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `SpaceMembership`. */
  spaceMemberships: SpaceMembershipsConnection;
  updatedAt: Scalars['Datetime'];
};


export type SpaceSpaceMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceMembershipCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** A condition to be used against `Space` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type SpaceCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `description` field. */
  description?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** An input for mutations affecting `Space` */
export type SpaceInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  description?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['UUID']>;
  name: Scalars['String'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

export type SpaceMembership = Node & {
  __typename?: 'SpaceMembership';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  personId: Scalars['UUID'];
  /** Reads and enables pagination through a set of `Post`. */
  postsByMembershipId: PostsConnection;
  roleId: Scalars['String'];
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  spaceId: Scalars['UUID'];
  updatedAt: Scalars['Datetime'];
};


export type SpaceMembershipPostsByMembershipIdArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PostCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};

/**
 * A condition to be used against `SpaceMembership` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type SpaceMembershipCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `roleId` field. */
  roleId?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `spaceId` field. */
  spaceId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** An input for mutations affecting `SpaceMembership` */
export type SpaceMembershipInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId: Scalars['UUID'];
  roleId: Scalars['String'];
  spaceId: Scalars['UUID'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `SpaceMembership`. Fields that are set will be updated. */
export type SpaceMembershipPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  roleId?: InputMaybe<Scalars['String']>;
  spaceId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `SpaceMembership` values. */
export type SpaceMembershipsConnection = {
  __typename?: 'SpaceMembershipsConnection';
  /** A list of edges which contains the `SpaceMembership` and cursor to aid in pagination. */
  edges: Array<SpaceMembershipsEdge>;
  /** A list of `SpaceMembership` objects. */
  nodes: Array<SpaceMembership>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `SpaceMembership` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `SpaceMembership` edge in the connection. */
export type SpaceMembershipsEdge = {
  __typename?: 'SpaceMembershipsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `SpaceMembership` at the end of the edge. */
  node: SpaceMembership;
};

/** Methods to use when ordering `SpaceMembership`. */
export enum SpaceMembershipsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  RoleIdAsc = 'ROLE_ID_ASC',
  RoleIdDesc = 'ROLE_ID_DESC',
  SpaceIdAsc = 'SPACE_ID_ASC',
  SpaceIdDesc = 'SPACE_ID_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** Represents an update to a `Space`. Fields that are set will be updated. */
export type SpacePatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  description?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `Space` values. */
export type SpacesConnection = {
  __typename?: 'SpacesConnection';
  /** A list of edges which contains the `Space` and cursor to aid in pagination. */
  edges: Array<SpacesEdge>;
  /** A list of `Space` objects. */
  nodes: Array<Space>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Space` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Space` edge in the connection. */
export type SpacesEdge = {
  __typename?: 'SpacesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Space` at the end of the edge. */
  node: Space;
};

/** Methods to use when ordering `Space`. */
export enum SpacesOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  DescriptionAsc = 'DESCRIPTION_ASC',
  DescriptionDesc = 'DESCRIPTION_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

export type Topic = Node & {
  __typename?: 'Topic';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** Reads and enables pagination through a set of `Interest`. */
  interests: InterestsConnection;
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  updatedAt: Scalars['Datetime'];
};


export type TopicInterestsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<InterestCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** A condition to be used against `Topic` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type TopicCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** An input for mutations affecting `Topic` */
export type TopicInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name: Scalars['String'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Topic`. Fields that are set will be updated. */
export type TopicPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `Topic` values. */
export type TopicsConnection = {
  __typename?: 'TopicsConnection';
  /** A list of edges which contains the `Topic` and cursor to aid in pagination. */
  edges: Array<TopicsEdge>;
  /** A list of `Topic` objects. */
  nodes: Array<Topic>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Topic` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Topic` edge in the connection. */
export type TopicsEdge = {
  __typename?: 'TopicsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Topic` at the end of the edge. */
  node: Topic;
};

/** Methods to use when ordering `Topic`. */
export enum TopicsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** All input for the `updateFamilyMembershipByNodeId` mutation. */
export type UpdateFamilyMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `FamilyMembership` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `FamilyMembership` being updated. */
  patch: FamilyMembershipPatch;
};

/** All input for the `updateFamilyMembership` mutation. */
export type UpdateFamilyMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `FamilyMembership` being updated. */
  patch: FamilyMembershipPatch;
};

/** The output of our update `FamilyMembership` mutation. */
export type UpdateFamilyMembershipPayload = {
  __typename?: 'UpdateFamilyMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  /** The `FamilyMembership` that was updated by this mutation. */
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our update `FamilyMembership` mutation. */
export type UpdateFamilyMembershipPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the `updateInterestByNodeId` mutation. */
export type UpdateInterestByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Interest` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Interest` being updated. */
  patch: InterestPatch;
};

/** All input for the `updateInterest` mutation. */
export type UpdateInterestInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Interest` being updated. */
  patch: InterestPatch;
};

/** The output of our update `Interest` mutation. */
export type UpdateInterestPayload = {
  __typename?: 'UpdateInterestPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Interest` that was updated by this mutation. */
  interest?: Maybe<Interest>;
  /** An edge for our `Interest`. May be used by Relay 1. */
  interestEdge?: Maybe<InterestsEdge>;
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
};


/** The output of our update `Interest` mutation. */
export type UpdateInterestPayloadInterestEdgeArgs = {
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** All input for the `updatePersonByNodeId` mutation. */
export type UpdatePersonByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Person` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Person` being updated. */
  patch: PersonPatch;
};

/** All input for the `updatePerson` mutation. */
export type UpdatePersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Person` being updated. */
  patch: PersonPatch;
};

/** The output of our update `Person` mutation. */
export type UpdatePersonPayload = {
  __typename?: 'UpdatePersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Person` that was updated by this mutation. */
  person?: Maybe<Person>;
  /** An edge for our `Person`. May be used by Relay 1. */
  personEdge?: Maybe<PeopleEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our update `Person` mutation. */
export type UpdatePersonPayloadPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};

/** All input for the `updatePostByNodeId` mutation. */
export type UpdatePostByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Post` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Post` being updated. */
  patch: PostPatch;
};

/** All input for the `updatePost` mutation. */
export type UpdatePostInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Post` being updated. */
  patch: PostPatch;
};

/** The output of our update `Post` mutation. */
export type UpdatePostPayload = {
  __typename?: 'UpdatePostPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `SpaceMembership` that is related to this `Post`. */
  membership?: Maybe<SpaceMembership>;
  /** The `Post` that was updated by this mutation. */
  post?: Maybe<Post>;
  /** An edge for our `Post`. May be used by Relay 1. */
  postEdge?: Maybe<PostsEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our update `Post` mutation. */
export type UpdatePostPayloadPostEdgeArgs = {
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};

/** All input for the `updateSpaceByNodeId` mutation. */
export type UpdateSpaceByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Space` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Space` being updated. */
  patch: SpacePatch;
};

/** All input for the `updateSpace` mutation. */
export type UpdateSpaceInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Space` being updated. */
  patch: SpacePatch;
};

/** All input for the `updateSpaceMembershipByNodeId` mutation. */
export type UpdateSpaceMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `SpaceMembership` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `SpaceMembership` being updated. */
  patch: SpaceMembershipPatch;
};

/** All input for the `updateSpaceMembership` mutation. */
export type UpdateSpaceMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `SpaceMembership` being updated. */
  patch: SpaceMembershipPatch;
};

/** The output of our update `SpaceMembership` mutation. */
export type UpdateSpaceMembershipPayload = {
  __typename?: 'UpdateSpaceMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  /** The `SpaceMembership` that was updated by this mutation. */
  spaceMembership?: Maybe<SpaceMembership>;
  /** An edge for our `SpaceMembership`. May be used by Relay 1. */
  spaceMembershipEdge?: Maybe<SpaceMembershipsEdge>;
};


/** The output of our update `SpaceMembership` mutation. */
export type UpdateSpaceMembershipPayloadSpaceMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** The output of our update `Space` mutation. */
export type UpdateSpacePayload = {
  __typename?: 'UpdateSpacePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Space` that was updated by this mutation. */
  space?: Maybe<Space>;
  /** An edge for our `Space`. May be used by Relay 1. */
  spaceEdge?: Maybe<SpacesEdge>;
};


/** The output of our update `Space` mutation. */
export type UpdateSpacePayloadSpaceEdgeArgs = {
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};

/** All input for the `updateTopicByNodeId` mutation. */
export type UpdateTopicByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Topic` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Topic` being updated. */
  patch: TopicPatch;
};

/** All input for the `updateTopic` mutation. */
export type UpdateTopicInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Topic` being updated. */
  patch: TopicPatch;
};

/** The output of our update `Topic` mutation. */
export type UpdateTopicPayload = {
  __typename?: 'UpdateTopicPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Topic` that was updated by this mutation. */
  topic?: Maybe<Topic>;
  /** An edge for our `Topic`. May be used by Relay 1. */
  topicEdge?: Maybe<TopicsEdge>;
};


/** The output of our update `Topic` mutation. */
export type UpdateTopicPayloadTopicEdgeArgs = {
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};

/** All input for the `updateUserByNodeId` mutation. */
export type UpdateUserByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `User` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `User` being updated. */
  patch: UserPatch;
};

/** All input for the `updateUserByPersonId` mutation. */
export type UpdateUserByPersonIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** An object where the defined keys will be set on the `User` being updated. */
  patch: UserPatch;
  personId: Scalars['UUID'];
};

/** All input for the `updateUser` mutation. */
export type UpdateUserInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `User` being updated. */
  patch: UserPatch;
};

/** The output of our update `User` mutation. */
export type UpdateUserPayload = {
  __typename?: 'UpdateUserPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Person` that is related to this `User`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `User` that was updated by this mutation. */
  user?: Maybe<User>;
  /** An edge for our `User`. May be used by Relay 1. */
  userEdge?: Maybe<UsersEdge>;
};


/** The output of our update `User` mutation. */
export type UpdateUserPayloadUserEdgeArgs = {
  orderBy?: InputMaybe<Array<UsersOrderBy>>;
};

export type User = Node & {
  __typename?: 'User';
  /** Reads and enables pagination through a set of `Authentication`. */
  authentications: AuthenticationsConnection;
  avatarUrl?: Maybe<Scalars['String']>;
  createdAt: Scalars['Datetime'];
  /** Reads a single `Family` that is related to this `User`. */
  family?: Maybe<Family>;
  id: Scalars['UUID'];
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `User`. */
  person?: Maybe<Person>;
  personId?: Maybe<Scalars['UUID']>;
  updatedAt: Scalars['Datetime'];
};


export type UserAuthenticationsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<AuthenticationCondition>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<AuthenticationsOrderBy>>;
};

/** A condition to be used against `User` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type UserCondition = {
  /** Checks for equality with the object’s `avatarUrl` field. */
  avatarUrl?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `User`. Fields that are set will be updated. */
export type UserPatch = {
  avatarUrl?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  personId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `User` values. */
export type UsersConnection = {
  __typename?: 'UsersConnection';
  /** A list of edges which contains the `User` and cursor to aid in pagination. */
  edges: Array<UsersEdge>;
  /** A list of `User` objects. */
  nodes: Array<User>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `User` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `User` edge in the connection. */
export type UsersEdge = {
  __typename?: 'UsersEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `User` at the end of the edge. */
  node: User;
};

/** Methods to use when ordering `User`. */
export enum UsersOrderBy {
  AvatarUrlAsc = 'AVATAR_URL_ASC',
  AvatarUrlDesc = 'AVATAR_URL_DESC',
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

export type AllSpacesQueryVariables = Exact<{ [key: string]: never; }>;


export type AllSpacesQuery = { __typename?: 'Query', spaces?: { __typename?: 'SpacesConnection', edges: Array<{ __typename?: 'SpacesEdge', node: { __typename?: 'Space', id: any, name: string } }> } | null };

export type CreateNewFamilyMemberMutationVariables = Exact<{
  name: Scalars['String'];
  role: Scalars['String'];
}>;


export type CreateNewFamilyMemberMutation = { __typename?: 'Mutation', createNewFamilyMember?: { __typename?: 'CreateNewFamilyMemberPayload', clientMutationId?: string | null } | null };

export type CreateSpaceMutationVariables = Exact<{
  name: Scalars['String'];
}>;


export type CreateSpaceMutation = { __typename?: 'Mutation', createSpace?: { __typename?: 'CreateSpacePayload', clientMutationId?: string | null } | null };

export type CurrentUserQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserQuery = { __typename?: 'Query', currentUser?: { __typename?: 'User', id: any, name: string } | null };

export type CurrentUserFamilyQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserFamilyQuery = { __typename?: 'Query', currentUser?: { __typename?: 'User', id: any, name: string, family?: { __typename?: 'Family', id: any, familyMemberships: { __typename?: 'FamilyMembershipsConnection', nodes: Array<{ __typename?: 'FamilyMembership', id: any, role: string, person?: { __typename?: 'Person', id: any, name: string } | null }> } } | null } | null };

export type CurrentFamilyMembershipQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentFamilyMembershipQuery = { __typename?: 'Query', currentFamilyMembership?: { __typename?: 'FamilyMembership', id: any, role: string, family?: { __typename?: 'Family', id: any } | null, person?: { __typename?: 'Person', id: any, name: string, user?: { __typename?: 'User', id: any } | null } | null } | null };

export type SpaceQueryVariables = Exact<{
  id: Scalars['UUID'];
}>;


export type SpaceQuery = { __typename?: 'Query', space?: { __typename?: 'Space', id: any, name: string, description?: string | null } | null };


export const AllSpacesDocument = gql`
    query AllSpaces {
  spaces {
    edges {
      node {
        id
        name
      }
    }
  }
}
    `;

/**
 * __useAllSpacesQuery__
 *
 * To run a query within a React component, call `useAllSpacesQuery` and pass it any options that fit your needs.
 * When your component renders, `useAllSpacesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAllSpacesQuery({
 *   variables: {
 *   },
 * });
 */
export function useAllSpacesQuery(baseOptions?: Apollo.QueryHookOptions<AllSpacesQuery, AllSpacesQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<AllSpacesQuery, AllSpacesQueryVariables>(AllSpacesDocument, options);
      }
export function useAllSpacesLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AllSpacesQuery, AllSpacesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<AllSpacesQuery, AllSpacesQueryVariables>(AllSpacesDocument, options);
        }
export type AllSpacesQueryHookResult = ReturnType<typeof useAllSpacesQuery>;
export type AllSpacesLazyQueryHookResult = ReturnType<typeof useAllSpacesLazyQuery>;
export type AllSpacesQueryResult = Apollo.QueryResult<AllSpacesQuery, AllSpacesQueryVariables>;
export const CreateNewFamilyMemberDocument = gql`
    mutation CreateNewFamilyMember($name: String!, $role: String!) {
  createNewFamilyMember(input: {name: $name, role: $role}) {
    clientMutationId
  }
}
    `;
export type CreateNewFamilyMemberMutationFn = Apollo.MutationFunction<CreateNewFamilyMemberMutation, CreateNewFamilyMemberMutationVariables>;

/**
 * __useCreateNewFamilyMemberMutation__
 *
 * To run a mutation, you first call `useCreateNewFamilyMemberMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateNewFamilyMemberMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createNewFamilyMemberMutation, { data, loading, error }] = useCreateNewFamilyMemberMutation({
 *   variables: {
 *      name: // value for 'name'
 *      role: // value for 'role'
 *   },
 * });
 */
export function useCreateNewFamilyMemberMutation(baseOptions?: Apollo.MutationHookOptions<CreateNewFamilyMemberMutation, CreateNewFamilyMemberMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateNewFamilyMemberMutation, CreateNewFamilyMemberMutationVariables>(CreateNewFamilyMemberDocument, options);
      }
export type CreateNewFamilyMemberMutationHookResult = ReturnType<typeof useCreateNewFamilyMemberMutation>;
export type CreateNewFamilyMemberMutationResult = Apollo.MutationResult<CreateNewFamilyMemberMutation>;
export type CreateNewFamilyMemberMutationOptions = Apollo.BaseMutationOptions<CreateNewFamilyMemberMutation, CreateNewFamilyMemberMutationVariables>;
export const CreateSpaceDocument = gql`
    mutation CreateSpace($name: String!) {
  createSpace(input: {space: {name: $name}}) {
    clientMutationId
  }
}
    `;
export type CreateSpaceMutationFn = Apollo.MutationFunction<CreateSpaceMutation, CreateSpaceMutationVariables>;

/**
 * __useCreateSpaceMutation__
 *
 * To run a mutation, you first call `useCreateSpaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateSpaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createSpaceMutation, { data, loading, error }] = useCreateSpaceMutation({
 *   variables: {
 *      name: // value for 'name'
 *   },
 * });
 */
export function useCreateSpaceMutation(baseOptions?: Apollo.MutationHookOptions<CreateSpaceMutation, CreateSpaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateSpaceMutation, CreateSpaceMutationVariables>(CreateSpaceDocument, options);
      }
export type CreateSpaceMutationHookResult = ReturnType<typeof useCreateSpaceMutation>;
export type CreateSpaceMutationResult = Apollo.MutationResult<CreateSpaceMutation>;
export type CreateSpaceMutationOptions = Apollo.BaseMutationOptions<CreateSpaceMutation, CreateSpaceMutationVariables>;
export const CurrentUserDocument = gql`
    query CurrentUser {
  currentUser {
    id
    name
  }
}
    `;

/**
 * __useCurrentUserQuery__
 *
 * To run a query within a React component, call `useCurrentUserQuery` and pass it any options that fit your needs.
 * When your component renders, `useCurrentUserQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCurrentUserQuery({
 *   variables: {
 *   },
 * });
 */
export function useCurrentUserQuery(baseOptions?: Apollo.QueryHookOptions<CurrentUserQuery, CurrentUserQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<CurrentUserQuery, CurrentUserQueryVariables>(CurrentUserDocument, options);
      }
export function useCurrentUserLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CurrentUserQuery, CurrentUserQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<CurrentUserQuery, CurrentUserQueryVariables>(CurrentUserDocument, options);
        }
export type CurrentUserQueryHookResult = ReturnType<typeof useCurrentUserQuery>;
export type CurrentUserLazyQueryHookResult = ReturnType<typeof useCurrentUserLazyQuery>;
export type CurrentUserQueryResult = Apollo.QueryResult<CurrentUserQuery, CurrentUserQueryVariables>;
export const CurrentUserFamilyDocument = gql`
    query CurrentUserFamily {
  currentUser {
    id
    name
    family {
      id
      familyMemberships {
        nodes {
          id
          person {
            id
            name
          }
          role
        }
      }
    }
  }
}
    `;

/**
 * __useCurrentUserFamilyQuery__
 *
 * To run a query within a React component, call `useCurrentUserFamilyQuery` and pass it any options that fit your needs.
 * When your component renders, `useCurrentUserFamilyQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCurrentUserFamilyQuery({
 *   variables: {
 *   },
 * });
 */
export function useCurrentUserFamilyQuery(baseOptions?: Apollo.QueryHookOptions<CurrentUserFamilyQuery, CurrentUserFamilyQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<CurrentUserFamilyQuery, CurrentUserFamilyQueryVariables>(CurrentUserFamilyDocument, options);
      }
export function useCurrentUserFamilyLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CurrentUserFamilyQuery, CurrentUserFamilyQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<CurrentUserFamilyQuery, CurrentUserFamilyQueryVariables>(CurrentUserFamilyDocument, options);
        }
export type CurrentUserFamilyQueryHookResult = ReturnType<typeof useCurrentUserFamilyQuery>;
export type CurrentUserFamilyLazyQueryHookResult = ReturnType<typeof useCurrentUserFamilyLazyQuery>;
export type CurrentUserFamilyQueryResult = Apollo.QueryResult<CurrentUserFamilyQuery, CurrentUserFamilyQueryVariables>;
export const CurrentFamilyMembershipDocument = gql`
    query CurrentFamilyMembership {
  currentFamilyMembership {
    id
    role
    family {
      id
    }
    person {
      id
      name
      user {
        id
      }
    }
  }
}
    `;

/**
 * __useCurrentFamilyMembershipQuery__
 *
 * To run a query within a React component, call `useCurrentFamilyMembershipQuery` and pass it any options that fit your needs.
 * When your component renders, `useCurrentFamilyMembershipQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCurrentFamilyMembershipQuery({
 *   variables: {
 *   },
 * });
 */
export function useCurrentFamilyMembershipQuery(baseOptions?: Apollo.QueryHookOptions<CurrentFamilyMembershipQuery, CurrentFamilyMembershipQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<CurrentFamilyMembershipQuery, CurrentFamilyMembershipQueryVariables>(CurrentFamilyMembershipDocument, options);
      }
export function useCurrentFamilyMembershipLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CurrentFamilyMembershipQuery, CurrentFamilyMembershipQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<CurrentFamilyMembershipQuery, CurrentFamilyMembershipQueryVariables>(CurrentFamilyMembershipDocument, options);
        }
export type CurrentFamilyMembershipQueryHookResult = ReturnType<typeof useCurrentFamilyMembershipQuery>;
export type CurrentFamilyMembershipLazyQueryHookResult = ReturnType<typeof useCurrentFamilyMembershipLazyQuery>;
export type CurrentFamilyMembershipQueryResult = Apollo.QueryResult<CurrentFamilyMembershipQuery, CurrentFamilyMembershipQueryVariables>;
export const SpaceDocument = gql`
    query Space($id: UUID!) {
  space(id: $id) {
    id
    name
    description
  }
}
    `;

/**
 * __useSpaceQuery__
 *
 * To run a query within a React component, call `useSpaceQuery` and pass it any options that fit your needs.
 * When your component renders, `useSpaceQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSpaceQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useSpaceQuery(baseOptions: Apollo.QueryHookOptions<SpaceQuery, SpaceQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<SpaceQuery, SpaceQueryVariables>(SpaceDocument, options);
      }
export function useSpaceLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<SpaceQuery, SpaceQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<SpaceQuery, SpaceQueryVariables>(SpaceDocument, options);
        }
export type SpaceQueryHookResult = ReturnType<typeof useSpaceQuery>;
export type SpaceLazyQueryHookResult = ReturnType<typeof useSpaceLazyQuery>;
export type SpaceQueryResult = Apollo.QueryResult<SpaceQuery, SpaceQueryVariables>;