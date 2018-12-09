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
		* db connect (MISS)
		* httpd serve
		* db disconnect (MISS)
	* stop
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
	* connect (MISS)
		* engine connect
		* if db exists check migrations
		* else create
	* disconnect (MISS)
		* engine disconnect
