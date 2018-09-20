(function(app) {
    var NgModule = ng.core.NgModule;
    var BrowserModule = ng.platformBrowser.BrowserModule;
    var CurrentTimeComponent = app.CurrentTimeComponent;
    var AppComponent = app.AppComponent;
    var TimeService = app.TimeService;
    
    app.AppModule = NgModule({
	   imports: [BrowserModule],
	   declarations: [AppComponent, CurrentTimeComponent],
       providers: [TimeService],
	   bootstrap: [AppComponent]
    })
    .Class({
	   constructor: function() {}
    });
    
})(window.app || (window.app = {}));
