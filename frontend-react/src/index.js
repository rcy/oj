import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import reportWebVitals from "./reportWebVitals";
import {
  HttpLink,
  ApolloClient,
  ApolloProvider,
  concat,
  InMemoryCache,
  split,
} from "@apollo/client";
import { getMainDefinition } from "@apollo/client/utilities";
import { setContext } from "@apollo/client/link/context";
import { GraphQLWsLink } from "@apollo/client/link/subscriptions";
import { createClient } from "graphql-ws";
import { ChakraProvider } from "@chakra-ui/react";

import { extendTheme } from "@chakra-ui/react";
import { mode } from "@chakra-ui/theme-tools";

const theme = extendTheme({
  styles: {
    global: (props) => ({
      body: {
        bg: mode(
          'linear-gradient(90deg, rgba(131,58,180,1) 0%, rgba(253,29,29,1) 50%, rgba(252,176,69,1) 100%);',
          'pink'
        )(props)
      }
    })
  }
})


//import { , concat } from "apollo-link";

console.log("env", process.env);

const httpLink = new HttpLink({ uri: "/graphql" });

const wsUrl =
  process.env.NODE_ENV === "development"
    ? "ws://localhost:5000/graphql"
    : "wss://octopusjr.ca/graphql";
const wsLink = new GraphQLWsLink(
  createClient({
    url: wsUrl,
  })
);

// cookies are used for user auth, but we add a header for current family membership here
// const familyMembershipMiddleware = new ApolloLink((operation, forward) => {
//   operation.setContext(() => ({
//     headers: {
//       'X-FAMILY-MEMBERSHIP-ID': JSON.parse(sessionStorage.getItem('familyMembershipId') || "null")
//     },
//   }));
//   return forward(operation);
// });

const asyncSettingsLink = setContext((_request) => {
  const skey = localStorage.getItem("sessionKey");
  const headers = skey
    ? {
        "x-person-session": skey,
      }
    : {};
  return new Promise((resolve, _reject) => {
    resolve({
      headers,
    });
  });
});

// The split function takes three parameters:
//
// * A function that's called for each operation to execute
// * The Link to use for an operation if the function returns a "truthy" value
// * The Link to use for an operation if the function returns a "falsy" value
const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === "OperationDefinition" &&
      definition.operation === "subscription"
    );
  },
  wsLink,
  concat(asyncSettingsLink, httpLink)
);

const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache(),
});

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <ChakraProvider theme={theme}>
        <App />
      </ChakraProvider>
    </ApolloProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
