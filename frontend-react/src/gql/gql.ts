/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

const documents = {
    "query AllSpaces($personId: UUID!) {\n  spaces {\n    totalCount\n    edges {\n      node {\n        id\n        name\n        spaceMemberships(condition: {personId: $personId}) {\n          edges {\n            node {\n              id\n              person {\n                id\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}": types.AllSpacesDocument,
    "mutation CreateNewFamilyMember($name: String!, $role: String!) {\n  createNewFamilyMember(input: {name: $name, role: $role}) {\n    familyMembership {\n      personId\n    }\n  }\n}": types.CreateNewFamilyMemberDocument,
    "mutation CreateSpace($name: String!) {\n  createSpace(input: {space: {name: $name}}) {\n    space {\n      id\n    }\n  }\n}": types.CreateSpaceDocument,
    "mutation CreateLoginCode($username: String!) {\n  createLoginCode(input: {username: $username}) {\n    loginCodeId\n  }\n}": types.CreateLoginCodeDocument,
    "query CurrentPerson {\n  currentPerson {\n    id\n    name\n    avatarUrl\n  }\n}": types.CurrentPersonDocument,
    "query CurrentUser {\n  currentUser {\n    id\n    name\n  }\n}": types.CurrentUserDocument,
    "query CurrentUserFamily {\n  currentUser {\n    id\n    name\n    family {\n      id\n      familyMemberships(orderBy: CREATED_AT_ASC) {\n        nodes {\n          id\n          title\n          person {\n            id\n            name\n            avatarUrl\n            user {\n              id\n            }\n          }\n          role\n        }\n      }\n    }\n  }\n}": types.CurrentUserFamilyDocument,
    "query CurrentFamilyMembership {\n  currentFamilyMembership {\n    id\n    role\n    title\n    family {\n      id\n    }\n    person {\n      id\n      name\n      avatarUrl\n      user {\n        id\n      }\n    }\n  }\n}": types.CurrentFamilyMembershipDocument,
    "query CurrentUserWithManagedPeople {\n  currentUser {\n    id\n    name\n    person {\n      id\n      name\n      avatarUrl\n    }\n    managedPeople {\n      nodes {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}": types.CurrentUserWithManagedPeopleDocument,
    "mutation ExchangeCode($loginCodeId: UUID!, $code: String!) {\n  exchangeCode(input: {loginCodeId: $loginCodeId, code: $code}) {\n    sessionKey\n  }\n}": types.ExchangeCodeDocument,
    "mutation JoinSpace($spaceId: UUID!, $personId: UUID!) {\n  createSpaceMembership(\n    input: {spaceMembership: {personId: $personId, spaceId: $spaceId, roleId: \"member\"}}\n  ) {\n    clientMutationId\n  }\n}": types.JoinSpaceDocument,
    "mutation PostMessage($membershipId: UUID!, $body: String!) {\n  postMessage(input: {spaceMembershipId: $membershipId, body: $body}) {\n    post {\n      id\n    }\n  }\n}": types.PostMessageDocument,
    "mutation SetPersonAvatar($personId: UUID!, $avatarUrl: String!) {\n  updatePerson(input: {id: $personId, patch: {avatarUrl: $avatarUrl}}) {\n    clientMutationId\n  }\n}": types.SetPersonAvatarDocument,
    "query Space($id: UUID!) {\n  space(id: $id) {\n    id\n    name\n    description\n  }\n}": types.SpaceDocument,
    "query SpaceMembershipByPersonIdAndSpaceId($personId: UUID!, $spaceId: UUID!) {\n  spaceMembershipByPersonIdAndSpaceId(personId: $personId, spaceId: $spaceId) {\n    id\n  }\n}": types.SpaceMembershipByPersonIdAndSpaceIdDocument,
    "query SpaceMembershipsByPersonId($personId: UUID!) {\n  spaceMemberships(condition: {personId: $personId}) {\n    edges {\n      node {\n        id\n        space {\n          id\n          name\n        }\n      }\n    }\n  }\n}": types.SpaceMembershipsByPersonIdDocument,
    "query SpaceMembershipsBySpaceId($spaceId: UUID!) {\n  spaceMemberships(condition: {spaceId: $spaceId}) {\n    edges {\n      node {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}": types.SpaceMembershipsBySpaceIdDocument,
    "query SpacePosts($spaceId: UUID!, $limit: Int) {\n  posts(condition: {spaceId: $spaceId}, last: $limit, orderBy: CREATED_AT_ASC) {\n    nodes {\n      id\n      body\n      membership {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}": types.SpacePostsDocument,
    "subscription SpacePostsAdded($spaceId: UUID!) {\n  posts(spaceId: $spaceId) {\n    event\n    post {\n      id\n      body\n      membership {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}": types.SpacePostsAddedDocument,
    "\n  fragment FamilyMembershipItem on FamilyMembership {\n    id\n    role\n    title\n    person {\n      id\n      name\n      avatarUrl\n      user {\n        id\n      }\n    }\n  }\n": types.FamilyMembershipItemFragmentDoc,
    "\nquery CurrentPersonFamilyMembership {\n  currentPerson {\n    id\n    familyMembership {\n      id\n      role\n      family {\n        id\n        familyMemberships(orderBy: CREATED_AT_ASC) {\n          edges {\n            node {\n              ...FamilyMembershipItem\n            }\n          }\n        }\n      }\n    }\n  }\n}\n": types.CurrentPersonFamilyMembershipDocument,
    "mutation CreateSpaceMembership($spaceId: UUID!, $personId: UUID!) {\n  createSpaceMembership(\n    input: {spaceMembership: {personId: $personId, spaceId: $spaceId, roleId: \"member\"}}\n  ) {\n    clientMutationId\n  }\n}": types.CreateSpaceMembershipDocument,
    "query SharedSpaces($person1: UUID!, $person2: UUID!) {\n  spaces(\n    filter: {and: [{spaceMemberships: {some: {personId: {equalTo: $person1}}}}, {spaceMemberships: {some: {personId: {equalTo: $person2}}}}]}\n  ) {\n    edges {\n      node {\n        id\n        name\n      }\n    }\n  }\n}": types.SharedSpacesDocument,
    "query PersonPageData($id: UUID!) {\n  person(id: $id) {\n    id\n    name\n    avatarUrl\n    createdAt\n  }\n}": types.PersonPageDataDocument,
};

export function graphql(source: "query AllSpaces($personId: UUID!) {\n  spaces {\n    totalCount\n    edges {\n      node {\n        id\n        name\n        spaceMemberships(condition: {personId: $personId}) {\n          edges {\n            node {\n              id\n              person {\n                id\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query AllSpaces($personId: UUID!) {\n  spaces {\n    totalCount\n    edges {\n      node {\n        id\n        name\n        spaceMemberships(condition: {personId: $personId}) {\n          edges {\n            node {\n              id\n              person {\n                id\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "mutation CreateNewFamilyMember($name: String!, $role: String!) {\n  createNewFamilyMember(input: {name: $name, role: $role}) {\n    familyMembership {\n      personId\n    }\n  }\n}"): (typeof documents)["mutation CreateNewFamilyMember($name: String!, $role: String!) {\n  createNewFamilyMember(input: {name: $name, role: $role}) {\n    familyMembership {\n      personId\n    }\n  }\n}"];
export function graphql(source: "mutation CreateSpace($name: String!) {\n  createSpace(input: {space: {name: $name}}) {\n    space {\n      id\n    }\n  }\n}"): (typeof documents)["mutation CreateSpace($name: String!) {\n  createSpace(input: {space: {name: $name}}) {\n    space {\n      id\n    }\n  }\n}"];
export function graphql(source: "mutation CreateLoginCode($username: String!) {\n  createLoginCode(input: {username: $username}) {\n    loginCodeId\n  }\n}"): (typeof documents)["mutation CreateLoginCode($username: String!) {\n  createLoginCode(input: {username: $username}) {\n    loginCodeId\n  }\n}"];
export function graphql(source: "query CurrentPerson {\n  currentPerson {\n    id\n    name\n    avatarUrl\n  }\n}"): (typeof documents)["query CurrentPerson {\n  currentPerson {\n    id\n    name\n    avatarUrl\n  }\n}"];
export function graphql(source: "query CurrentUser {\n  currentUser {\n    id\n    name\n  }\n}"): (typeof documents)["query CurrentUser {\n  currentUser {\n    id\n    name\n  }\n}"];
export function graphql(source: "query CurrentUserFamily {\n  currentUser {\n    id\n    name\n    family {\n      id\n      familyMemberships(orderBy: CREATED_AT_ASC) {\n        nodes {\n          id\n          title\n          person {\n            id\n            name\n            avatarUrl\n            user {\n              id\n            }\n          }\n          role\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query CurrentUserFamily {\n  currentUser {\n    id\n    name\n    family {\n      id\n      familyMemberships(orderBy: CREATED_AT_ASC) {\n        nodes {\n          id\n          title\n          person {\n            id\n            name\n            avatarUrl\n            user {\n              id\n            }\n          }\n          role\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "query CurrentFamilyMembership {\n  currentFamilyMembership {\n    id\n    role\n    title\n    family {\n      id\n    }\n    person {\n      id\n      name\n      avatarUrl\n      user {\n        id\n      }\n    }\n  }\n}"): (typeof documents)["query CurrentFamilyMembership {\n  currentFamilyMembership {\n    id\n    role\n    title\n    family {\n      id\n    }\n    person {\n      id\n      name\n      avatarUrl\n      user {\n        id\n      }\n    }\n  }\n}"];
export function graphql(source: "query CurrentUserWithManagedPeople {\n  currentUser {\n    id\n    name\n    person {\n      id\n      name\n      avatarUrl\n    }\n    managedPeople {\n      nodes {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query CurrentUserWithManagedPeople {\n  currentUser {\n    id\n    name\n    person {\n      id\n      name\n      avatarUrl\n    }\n    managedPeople {\n      nodes {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "mutation ExchangeCode($loginCodeId: UUID!, $code: String!) {\n  exchangeCode(input: {loginCodeId: $loginCodeId, code: $code}) {\n    sessionKey\n  }\n}"): (typeof documents)["mutation ExchangeCode($loginCodeId: UUID!, $code: String!) {\n  exchangeCode(input: {loginCodeId: $loginCodeId, code: $code}) {\n    sessionKey\n  }\n}"];
export function graphql(source: "mutation JoinSpace($spaceId: UUID!, $personId: UUID!) {\n  createSpaceMembership(\n    input: {spaceMembership: {personId: $personId, spaceId: $spaceId, roleId: \"member\"}}\n  ) {\n    clientMutationId\n  }\n}"): (typeof documents)["mutation JoinSpace($spaceId: UUID!, $personId: UUID!) {\n  createSpaceMembership(\n    input: {spaceMembership: {personId: $personId, spaceId: $spaceId, roleId: \"member\"}}\n  ) {\n    clientMutationId\n  }\n}"];
export function graphql(source: "mutation PostMessage($membershipId: UUID!, $body: String!) {\n  postMessage(input: {spaceMembershipId: $membershipId, body: $body}) {\n    post {\n      id\n    }\n  }\n}"): (typeof documents)["mutation PostMessage($membershipId: UUID!, $body: String!) {\n  postMessage(input: {spaceMembershipId: $membershipId, body: $body}) {\n    post {\n      id\n    }\n  }\n}"];
export function graphql(source: "mutation SetPersonAvatar($personId: UUID!, $avatarUrl: String!) {\n  updatePerson(input: {id: $personId, patch: {avatarUrl: $avatarUrl}}) {\n    clientMutationId\n  }\n}"): (typeof documents)["mutation SetPersonAvatar($personId: UUID!, $avatarUrl: String!) {\n  updatePerson(input: {id: $personId, patch: {avatarUrl: $avatarUrl}}) {\n    clientMutationId\n  }\n}"];
export function graphql(source: "query Space($id: UUID!) {\n  space(id: $id) {\n    id\n    name\n    description\n  }\n}"): (typeof documents)["query Space($id: UUID!) {\n  space(id: $id) {\n    id\n    name\n    description\n  }\n}"];
export function graphql(source: "query SpaceMembershipByPersonIdAndSpaceId($personId: UUID!, $spaceId: UUID!) {\n  spaceMembershipByPersonIdAndSpaceId(personId: $personId, spaceId: $spaceId) {\n    id\n  }\n}"): (typeof documents)["query SpaceMembershipByPersonIdAndSpaceId($personId: UUID!, $spaceId: UUID!) {\n  spaceMembershipByPersonIdAndSpaceId(personId: $personId, spaceId: $spaceId) {\n    id\n  }\n}"];
export function graphql(source: "query SpaceMembershipsByPersonId($personId: UUID!) {\n  spaceMemberships(condition: {personId: $personId}) {\n    edges {\n      node {\n        id\n        space {\n          id\n          name\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query SpaceMembershipsByPersonId($personId: UUID!) {\n  spaceMemberships(condition: {personId: $personId}) {\n    edges {\n      node {\n        id\n        space {\n          id\n          name\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "query SpaceMembershipsBySpaceId($spaceId: UUID!) {\n  spaceMemberships(condition: {spaceId: $spaceId}) {\n    edges {\n      node {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query SpaceMembershipsBySpaceId($spaceId: UUID!) {\n  spaceMemberships(condition: {spaceId: $spaceId}) {\n    edges {\n      node {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "query SpacePosts($spaceId: UUID!, $limit: Int) {\n  posts(condition: {spaceId: $spaceId}, last: $limit, orderBy: CREATED_AT_ASC) {\n    nodes {\n      id\n      body\n      membership {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query SpacePosts($spaceId: UUID!, $limit: Int) {\n  posts(condition: {spaceId: $spaceId}, last: $limit, orderBy: CREATED_AT_ASC) {\n    nodes {\n      id\n      body\n      membership {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "subscription SpacePostsAdded($spaceId: UUID!) {\n  posts(spaceId: $spaceId) {\n    event\n    post {\n      id\n      body\n      membership {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"): (typeof documents)["subscription SpacePostsAdded($spaceId: UUID!) {\n  posts(spaceId: $spaceId) {\n    event\n    post {\n      id\n      body\n      membership {\n        id\n        person {\n          id\n          name\n          avatarUrl\n        }\n      }\n    }\n  }\n}"];
export function graphql(source: "\n  fragment FamilyMembershipItem on FamilyMembership {\n    id\n    role\n    title\n    person {\n      id\n      name\n      avatarUrl\n      user {\n        id\n      }\n    }\n  }\n"): (typeof documents)["\n  fragment FamilyMembershipItem on FamilyMembership {\n    id\n    role\n    title\n    person {\n      id\n      name\n      avatarUrl\n      user {\n        id\n      }\n    }\n  }\n"];
export function graphql(source: "\nquery CurrentPersonFamilyMembership {\n  currentPerson {\n    id\n    familyMembership {\n      id\n      role\n      family {\n        id\n        familyMemberships(orderBy: CREATED_AT_ASC) {\n          edges {\n            node {\n              ...FamilyMembershipItem\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"): (typeof documents)["\nquery CurrentPersonFamilyMembership {\n  currentPerson {\n    id\n    familyMembership {\n      id\n      role\n      family {\n        id\n        familyMemberships(orderBy: CREATED_AT_ASC) {\n          edges {\n            node {\n              ...FamilyMembershipItem\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"];
export function graphql(source: "mutation CreateSpaceMembership($spaceId: UUID!, $personId: UUID!) {\n  createSpaceMembership(\n    input: {spaceMembership: {personId: $personId, spaceId: $spaceId, roleId: \"member\"}}\n  ) {\n    clientMutationId\n  }\n}"): (typeof documents)["mutation CreateSpaceMembership($spaceId: UUID!, $personId: UUID!) {\n  createSpaceMembership(\n    input: {spaceMembership: {personId: $personId, spaceId: $spaceId, roleId: \"member\"}}\n  ) {\n    clientMutationId\n  }\n}"];
export function graphql(source: "query SharedSpaces($person1: UUID!, $person2: UUID!) {\n  spaces(\n    filter: {and: [{spaceMemberships: {some: {personId: {equalTo: $person1}}}}, {spaceMemberships: {some: {personId: {equalTo: $person2}}}}]}\n  ) {\n    edges {\n      node {\n        id\n        name\n      }\n    }\n  }\n}"): (typeof documents)["query SharedSpaces($person1: UUID!, $person2: UUID!) {\n  spaces(\n    filter: {and: [{spaceMemberships: {some: {personId: {equalTo: $person1}}}}, {spaceMemberships: {some: {personId: {equalTo: $person2}}}}]}\n  ) {\n    edges {\n      node {\n        id\n        name\n      }\n    }\n  }\n}"];
export function graphql(source: "query PersonPageData($id: UUID!) {\n  person(id: $id) {\n    id\n    name\n    avatarUrl\n    createdAt\n  }\n}"): (typeof documents)["query PersonPageData($id: UUID!) {\n  person(id: $id) {\n    id\n    name\n    avatarUrl\n    createdAt\n  }\n}"];

export function graphql(source: string): unknown;
export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;