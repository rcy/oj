import dotenv from 'dotenv'; dotenv.config();
import createError from 'http-errors';
import express from 'express';
import path from 'path';
import cookieParser from 'cookie-parser';
import session from 'express-session';
import passport from 'passport';
import logger from 'morgan';
import { ensureLoggedIn } from 'connect-ensure-login';
import { Strategy as GoogleStrategy } from 'passport-google-oauth20';
import connectPgSimple from 'connect-pg-simple';
import { postgraphile } from 'postgraphile';
import pgSimplifyInflector from '@graphile-contrib/pg-simplify-inflector';

import pg from 'pg';
const pgPool = new pg.Pool({
  connectionString: process.env.DATABASE_URL,
});
pgPool.on('error', (err, client) => {
  console.error('Unexpected error on idle pg client', err);
  process.exit(-1);
});

passport.use(
  new GoogleStrategy({
    clientID: process.env.GOOGLE_CLIENT_ID,
    clientSecret: process.env.GOOGLE_CLIENT_SECRET,
    callbackURL: process.env.GOOGLE_CALLBACK_URL,
  }, async function(accessToken, refreshToken, profile, cb) {
    console.log('GoogleStrategy 1', profile)
    const client = await pgPool.connect();
    try {
      const res1 = await client.query("select user_id from app_public.authentications where service = 'google' and identifier = $1", [profile.id])
      console.log('GoogleStrategy 2.0', res1.rows[0])
      if (res1.rows[0]) {
        return cb(null, res1.rows[0].user_id)
      } else {
        const res2 = await client.query(`select app_public.create_user_authentication(service:='google', name:=$1, identifier:=$2, details:=$3)`,
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

app.use(logger('dev'));
app.use(express.json());
app.use(cookieParser());

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false, // don't save session if unmodified
  saveUninitialized: false, // don't create session until something stored
  store: new (connectPgSimple(session))({
    pool: pgPool,
    schemaName: 'app_private',
    tableName: 'passport_sessions',
  })
}));
app.use(passport.authenticate('session'));

app.get('/auth/google',
        passport.authenticate('google', {
          scope: ['profile', 'email'],
          prompt: 'select_account',
        })
);

app.get('/auth/google/callback',
        passport.authenticate('google', {
          failureRedirect: '/error',
          successRedirect: '/welcome',
        })
);

app.post('/auth/logout', function(req, res, next){
  req.logout(function(err) {
    if (err) { return next(err); }
    res.redirect('/');
  });
});

app.get('/', function(req,res) {
  res.send('<h1>welcome to aschool</h1> <a href="/auth/google">login</a>')
})

app.get('/welcome', ensureLoggedIn('/auth/google'), function(req,res) {
  res.send(`welcome ${req.user.id}`)
})

// https://www.graphile.org/postgraphile/usage-library/
app.use(postgraphile(pgPool, ['app_public'], {
  watchPg: true,
  ownerConnectionString: process.env.WATCH_DATABASE_URL,
  graphiql: true,
  enhanceGraphiql: true,
  //subscriptions: true,
  dynamicJson: true,
  setofFunctionsContainNulls: false,
  ignoreRBAC: false,
  showErrorStack: 'json',
  extendedErrors: ['hint', 'detail', 'errcode'],
  allowExplain: true, // don't use in production
  legacyRelations: 'omit',
  //exportGqlSchemaPath: `${__dirname}/schema.graphql`,
  sortExport: true,
  appendPlugins: [pgSimplifyInflector],
  pgSettings: function(req) {
    console.log('**************************************************************** req.user', req.user)
    return {
      'role': 'visitor',
      'user.id': req.user,
    }
  },
}));

console.log(`express listening on ${process.env.PORT}`)
app.listen(process.env.PORT)
