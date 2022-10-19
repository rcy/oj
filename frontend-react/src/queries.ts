import { graphql } from './gql';

export const CURRENT_USER = graphql(`
  query CurrentUser {
    currentUser {
      id
      name
    }
  }
`);

export const CURRENT_USER_FAMILY = graphql(`
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
`);

export const CURRENT_FAMILY_MEMBERSHIP = graphql(`
  query CurrentFamilyMembership{
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
`);

export const CREATE_NEW_FAMILY_MEMBER = graphql(`
  mutation CreateNewFamilyMember($name: String!, $role: String!) {
    createNewFamilyMember(input: {name: $name, role: $role}) {
      clientMutationId
    }
  }
`);

export const CREATE_SPACE = graphql(`
  mutation CreateSpace($name: String!) {
    createSpace(input: {space: {name: $name}}) {
      clientMutationId
    }
  }
`);

