(function(app) {
  var Component = ng.core.Component;
  var QuoteService = app.QuoteService;
  var Inject = ng.core.Inject;

  @Component({
    selector: 'random-quote',
    template: '<p><em>{{quote.line}}</em> - {{quote.author}}</p>'
  })
  class RandomQuoteComponent {
    constructor: (@Inject quoteService) {
        quoteService.generateRandomQuotes(2000, (quote) => this.quote = quote);
    }
  }

  app.RandomQuoteComponent = RandomQuoteComponent;

})(window.app || (window.app = {}));
