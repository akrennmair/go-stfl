typedef struct stfl_form stfl_form;

static struct stfl_ipool * ipool = NULL;

static void flush(void) {
	if (!ipool) {
		ipool = stfl_ipool_create("utf-8");
	}
	stfl_ipool_flush(ipool);
}

/* TODO: I can haz init routine where I can put an atexit(destroy)?
static void destroy(void) {
	stfl_ipool_destroy(ipool);
	ipool = NULL;
}
*/

#define TOWC(x) (stfl_ipool_towc(ipool, (x)))
#define FROMWC(x) (stfl_ipool_fromwc(ipool, (x)))

static stfl_form * stfl_create_wrapper(const char * text) {
	flush();
	return stfl_create(TOWC(text));
}

static const char * stfl_run_wrapper(stfl_form * form, int timeout) {
	flush();
	return FROMWC(stfl_run(form, timeout));
}

static const char * stfl_get_wrapper(stfl_form * form, const char * name) {
	flush();
	return FROMWC(stfl_get(form, TOWC(name)));
}

static void stfl_set_wrapper(stfl_form * form, const char * name, const char * value) {
	flush();
	stfl_set(form, TOWC(name), TOWC(value));
}

static const char * stfl_get_focus_wrapper(stfl_form * form) {
	flush();
	return FROMWC(stfl_get_focus(form));
}

static void stfl_set_focus_wrapper(stfl_form * form, const char * name) {
	flush();
	stfl_set_focus(form, TOWC(name));
}

static const char * stfl_quote_wrapper(const char * text) {
	flush();
	return FROMWC(stfl_quote(TOWC(text)));
}

static void stfl_modify_wrapper(stfl_form * form, const char * name, const char * mode, const char * text) {
	flush();
	stfl_modify(form, TOWC(name), TOWC(mode), TOWC(text));
}

static const char * stfl_lookup_wrapper(stfl_form * form, const char * path, const char * newname) {
	flush();
	return FROMWC(stfl_lookup(form, TOWC(path), TOWC(newname)));
}

static const char * stfl_dump_wrapper(stfl_form * form, const char * name, const char * prefix, int focus) {
	flush();
	return FROMWC(stfl_dump(form, TOWC(name), TOWC(prefix), focus));
}
