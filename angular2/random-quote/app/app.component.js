'use strict';

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

(function (app) {
  var _dec, _class;

  var Component = ng.core.Component;

  var AppComponent = (_dec = Component({
    selector: 'my-app',
    template: '\n      <h1>Random Quote!!!</h1>\n      <random-quote></random-quote>\n    '
  }), _dec(_class = function AppComponent() {
    _classCallCheck(this, AppComponent);
  }) || _class);


  app.AppComponent = AppComponent;
})(window.app || (window.app = {}));