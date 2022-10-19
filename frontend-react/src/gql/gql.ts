/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

const documents = {
    "\n  query CurrentUser {\n    currentUser {\n      id\n      name\n    }\n  }\n": types.CurrentUserDocument,
    "\n  query CurrentUserFamily {\n    currentUser {\n      id\n      name\n      family {\n        id\n        familyMemberships {\n          nodes {\n            id\n            person {\n              id\n              name\n            }\n            role\n          }\n        }\n      }\n    }\n  }\n": types.CurrentUserFamilyDocument,
    "\n  query CurrentFamilyMembership{\n    currentFamilyMembership {\n      id\n      role\n      family {\n        id\n      }\n      person {\n        id\n        name\n        user {\n          id\n        }\n      }\n    }\n  }\n": types.CurrentFamilyMembershipDocument,
    "\n  mutation CreateNewFamilyMember($name: String!, $role: String!) {\n    createNewFamilyMember(input: {name: $name, role: $role}) {\n      clientMutationId\n    }\n  }\n": types.CreateNewFamilyMemberDocument,
    "\n  mutation CreateSpace($name: String!) {\n    createSpace(input: {space: {name: $name}}) {\n      clientMutationId\n    }\n  }\n": types.CreateSpaceDocument,
    "\n  query AllSpaces {\n    spaces {\n      edges {\n        node {\n          id\n          name\n        }\n      }\n    }\n  }\n": types.AllSpacesDocument,
};

export function graphql(source: "\n  query CurrentUser {\n    currentUser {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  query CurrentUser {\n    currentUser {\n      id\n      name\n    }\n  }\n"];
export function graphql(source: "\n  query CurrentUserFamily {\n    currentUser {\n      id\n      name\n      family {\n        id\n        familyMemberships {\n          nodes {\n            id\n            person {\n              id\n              name\n            }\n            role\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query CurrentUserFamily {\n    currentUser {\n      id\n      name\n      family {\n        id\n        familyMemberships {\n          nodes {\n            id\n            person {\n              id\n              name\n            }\n            role\n          }\n        }\n      }\n    }\n  }\n"];
export function graphql(source: "\n  query CurrentFamilyMembership{\n    currentFamilyMembership {\n      id\n      role\n      family {\n        id\n      }\n      person {\n        id\n        name\n        user {\n          id\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query CurrentFamilyMembership{\n    currentFamilyMembership {\n      id\n      role\n      family {\n        id\n      }\n      person {\n        id\n        name\n        user {\n          id\n        }\n      }\n    }\n  }\n"];
export function graphql(source: "\n  mutation CreateNewFamilyMember($name: String!, $role: String!) {\n    createNewFamilyMember(input: {name: $name, role: $role}) {\n      clientMutationId\n    }\n  }\n"): (typeof documents)["\n  mutation CreateNewFamilyMember($name: String!, $role: String!) {\n    createNewFamilyMember(input: {name: $name, role: $role}) {\n      clientMutationId\n    }\n  }\n"];
export function graphql(source: "\n  mutation CreateSpace($name: String!) {\n    createSpace(input: {space: {name: $name}}) {\n      clientMutationId\n    }\n  }\n"): (typeof documents)["\n  mutation CreateSpace($name: String!) {\n    createSpace(input: {space: {name: $name}}) {\n      clientMutationId\n    }\n  }\n"];
export function graphql(source: "\n  query AllSpaces {\n    spaces {\n      edges {\n        node {\n          id\n          name\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query AllSpaces {\n    spaces {\n      edges {\n        node {\n          id\n          name\n        }\n      }\n    }\n  }\n"];

export function graphql(source: string): unknown;
export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;