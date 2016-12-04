'use strict';

(function (app) {
  var Component = ng.core.Component;

  app.AppComponent = Component({
    selector: 'my-app',
    template: '\n      <h1>Random Quote!!!</h1>\n      <random-quote></random-quote>\n    '
  }).Class({
    constructor: function AppComponent() {}
  });
})(window.app || (window.app = {}));