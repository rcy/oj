dev:
	npx nodemon server.js

react:
	cd ../frontend-react && BUILD_PATH=../backend/react npm run build

clean:
	rm -rf react

deploy: clean react
	flyctl deploy

# migrate targets
init:
	npx graphile-migrate init

watch:
	npm run watch

commit:
	npx graphile-migrate commit

uncommit:
	npx graphile-migrate uncommit

reset:
	npx graphile-migrate reset --erase
	${MAKE} create

hardreset: down
	rm -rf .gmrc migrations
