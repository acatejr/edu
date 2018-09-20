(function(app) {
    var Component = ng.core.Component;
    
    app.AppComponent = Component({
        selector: 'my-app',
        template: '<h1>Hello World!</h1>' + 
                  '<random-quote></random-quote>'
    })
    .Class({
        constructor: function AppComponent() {}
    });
    
})(window.app || (window.app = {}));
