import { gql } from '@apollo/client';

export const CURRENT_USER = gql`
  query CurrentUser {
    currentUser {
      id
      name
    }
  }
`;

export const CURRENT_USER_FAMILY = gql`
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

export const CURRENT_FAMILY_MEMBERSHIP = gql`
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
`;

export const CREATE_NEW_FAMILY_MEMBER = gql`
  mutation CreateNewFamilyMember($name: String!, $role: String!) {
    createNewFamilyMember(input: {name: $name, role: $role}) {
      clientMutationId
    }
  }
`
