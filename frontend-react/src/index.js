import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { HttpLink, ApolloLink, ApolloClient, ApolloProvider, concat, InMemoryCache } from '@apollo/client';
//import { , concat } from "apollo-link";

const httpLink = new HttpLink({ uri: '/graphql' });

// cookies are used for user auth, but we add a header for current family membership here
const familyMembershipMiddleware = new ApolloLink((operation, forward) => {
  operation.setContext({
    headers: {
      'X-FAMILY-MEMBERSHIP-ID': sessionStorage.getItem('familyMembershipId')
    },
  });
  return forward(operation);
});

const client = new ApolloClient({
  link: concat(familyMembershipMiddleware, httpLink),
  cache: new InMemoryCache(),
});

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
