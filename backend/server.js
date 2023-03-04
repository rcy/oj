import dotenv from 'dotenv'; dotenv.config({ path: '../.env' });
import express from 'express';
import cookieParser from 'cookie-parser';
import session from 'express-session';
import passport from 'passport';
import logger from 'morgan';
import { Strategy as GoogleStrategy } from 'passport-google-oauth20';
import pg from 'pg';
import connectPgSimple from 'connect-pg-simple';
import { postgraphile } from 'postgraphile';
import pgSimplifyInflector from '@graphile-contrib/pg-simplify-inflector';
import ConnectionFilterPlugin from "postgraphile-plugin-connection-filter";
import { worker } from './worker.js'

import { fileURLToPath } from 'url'
import { dirname } from 'path'

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

const pgAppPool = new pg.Pool({ connectionString: process.env.DATABASE_URL });
pgAppPool.on('error', (err, client) => {
  console.error('Unexpected error on idle pg client (pgAppPool)', err);
  process.exit(-1);
});

const pgVisitorPool = new pg.Pool({ connectionString: process.env.VISITOR_DATABASE_URL });
pgVisitorPool.on('error', (err, client) => {
  console.error('Unexpected error on idle pg client (pgVisitorPool)', err);
  process.exit(-1);
});

passport.use(
  new GoogleStrategy({
    clientID: process.env.GOOGLE_CLIENT_ID,
    clientSecret: process.env.GOOGLE_CLIENT_SECRET,
    callbackURL: process.env.BASE_URL + '/auth/google/callback',
  }, async function(accessToken, refreshToken, profile, cb) {
    console.log('GoogleStrategy 1', profile)
    const client = await pgAppPool.connect();

    // revisit this logic... we should maybe always create a new
    // authentication, but from there, look up the user by email rather
    // than creating a new user.  The result should be, we can delete
    // all the authentications, and have seamlessly continue to work
    // after users re-authenticate.  Currently if we delete
    // authentications, subsequent logins will result in new users and
    // families being created, orphaning the old ones

    try {
      const res1 = await client.query("select user_id from app_public.authentications where service = 'google' and identifier = $1", [profile.id])
      console.log('GoogleStrategy 2.0', res1.rows[0])
      if (res1.rows[0]) {
        return cb(null, res1.rows[0].user_id)
      } else {
        const res2 = await client.query(`select app_private.create_user_authentication(service:='google', name:=$1, identifier:=$2, details:=$3)`,
                                        [profile._json.name, profile.id, profile])
        return cb(null, res2.rows[0]['create_user_authentication'])
      }
    } finally {
      client.release()
    }
  }
));

passport.serializeUser(function(user, cb) {
  process.nextTick(function() {
    cb(null, user); //{ id: user.id, username: user.username, name: user.name });
  });
});

passport.deserializeUser(function(user, cb) {
  process.nextTick(function() {
    return cb(null, user);
  });
});

const app = express();

const environment = app.get('env');

app.use(logger('dev'));
app.use(express.json());
app.use(cookieParser());

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false, // don't save session if unmodified
  saveUninitialized: false, // don't create session until something stored
  store: new (connectPgSimple(session))({
    pool: pgAppPool,
    schemaName: 'app_private',
    tableName: 'passport_sessions',
  })
}));
app.use(passport.authenticate('session'));

app.get('/auth/login',
        function(req, res, next) {
          res.cookie('from', req.query.from)
          next()
        },
        passport.authenticate('google', {
          scope: ['profile', 'email'],
          //prompt: 'select_account',
        })
);

app.get('/auth/google/callback',
        passport.authenticate('google', {
          failureRedirect: '/error',
          successRedirect: '/auth/success',
        })
);

app.get('/auth/success', function(req, res, next) {
  res.redirect(req.cookies.from || '/welcome')
});

app.get('/auth/logout', function(req, res, next){
  req.logout(function(err) {
    if (err) { return next(err); }
    res.redirect('/');
  });
});

const delay = 0;

// https://www.graphile.org/postgraphile/usage-library/
const postgraphileOptions = Object.assign({
  //subscriptions: true,
  dynamicJson: true,
  setofFunctionsContainNulls: false,
  ignoreRBAC: false,
  showErrorStack: 'json',
  extendedErrors: ['hint', 'detail', 'errcode'],
  legacyRelations: 'omit',
  sortExport: true,
  appendPlugins: [pgSimplifyInflector, ConnectionFilterPlugin],
  graphileBuildOptions: {
    connectionFilterRelations: true,
  },
  pgSettings: async function (req) {
    const personId = req.headers['x-person-id'];
    console.log('**************************************************************** user/person', req.user, personId)

    if (delay) {
      await new Promise(function (resolve) {
        setTimeout(resolve, delay);
      });
    }

    return {
      'role': 'visitor',
      'user.id': req.user,
      'person.id': personId,
    }
  },
}, environment === 'development' && {
  watchPg: true,
  ownerConnectionString: process.env.WATCH_DATABASE_URL,
  graphiql: true,
  enhanceGraphiql: true,
  allowExplain: true,
  exportGqlSchemaPath: `${__dirname}/schema.graphql`,
})

app.use(postgraphile(pgVisitorPool, ['app_public'], postgraphileOptions));

// serve all other routes from react app
app.use(express.static(__dirname + '/react'));
app.get('*', (_, res) => res.sendFile(__dirname + '/react/index.html'));

console.log(`express listening on ${process.env.PORT}`)
app.listen(process.env.PORT)

worker().catch((err) => {
  console.error(err)
  process.exit(1)
})
