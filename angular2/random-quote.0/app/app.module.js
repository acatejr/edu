(function(app) {
    var NgModule = ng.core.NgModule;
    var BrowserModule = ng.platformBrowser.BrowserModule;
    // var platformBrowserDynamic = ng.platformBrowserDynamic.platformBrowserDynamic;
    var QuoteService = app.QuoteService;
    var RandomQuoteComponent = app.RandomQuoteComponent;
    var AppComponent = app.AppComponent;
    
    app.AppModule = NgModule({
	   imports: [BrowserModule],
	   declarations: [AppComponent, RandomQuoteComponent],
       providers: [QuoteService],
	   bootstrap: [AppComponent]
    })
    .Class({
	   constructor: function() {}
    });

    // platformBrowserDynamic().bootstrapModule(AppModule);
    
})(window.app || (window.app = {}));
