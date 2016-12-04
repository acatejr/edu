(function(app){
    var Class = ng.core.Class;

    app.TimeService = Class({
        constructor: function TimeService() {
            this.currentTime = Date();
        },
        getCurrentTime: function() {
            return Date();
        },
        generateCurrentTime: function(delay, callback) {
            var self = this;
            callback(this.getCurrentTime());
            setInterval(function() {
                callback(self.getCurrentTime())
            }, delay)
        }
    });

})(window.app || (window.app = {}));