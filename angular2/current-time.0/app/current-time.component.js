(function(app) {

    var Component = ng.core.Component;
    var TimeService = app.TimeService;

    app.CurrentTimeComponent = Component({
        selector: 'current-time',
        template: '<p><em>The current time is</em></p>' +
                  '<div>{{currentTime}}</div>'
    })
    .Class({
        constructor: [TimeService, function CurrentTimeComponent(timeService) {
            var self = this;
            timeService.generateCurrentTime(1000, function(ct) {
                self.currentTime = ct;
            });
        }]
    });
    
})(window.app || (window.app = {}));
