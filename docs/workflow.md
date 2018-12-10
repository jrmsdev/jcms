* jcms
	* main
		* parse cmd flags
		* webapp config
		* create webapp
		* start
		* serve
		* stop
	* start
		* webapp config defaults
		* log init
		* webapp setup
		* httpd setup
		* httpd listen
	* serve
		* httpd serve
		* db connect or serve error (MISS)
		* db schema check (MISS)
	* stop
		* db disconnect or serve error (?) (MISS)
		* httpd stop

## internal

* webapp
	* setup
		* setup assets manager
		* setup storage driver
		* setup db engine

* httpd
	* setup
		* init httpd router
		* setup httpd handlers
		* create httpd server

* db
	* setup
		* parse config database uri
		* load engine
	* connect
		* engine connect
	* disconnect (MISS)
		* engine disconnect

* db/schema
	* check (MISS)
		* if db exists check migrations
		* else create
