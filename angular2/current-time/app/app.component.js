(function(app) {
    var Component = ng.core.Component;
    
    app.AppComponent = Component({
        selector: 'app',
        template: '<h1>The current time is: </h1>' + 
                  "<current-time></current-time>"
    })
    .Class({
        constructor: function AppComponent() {}
    });
    
})(window.app || (window.app = {}));
