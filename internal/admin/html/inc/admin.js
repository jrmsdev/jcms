w3.includeHTML();

w3.getHttpObject("/_/jcms.json", jcmsDisplay);
function jcmsDisplay(obj) {
	w3.displayObject("jcms", obj);
}
