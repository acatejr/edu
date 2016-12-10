import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { TimeService } from './time.service';
import { CurrentTimeComponent } from './current-time.component';
import { AppComponent } from './app.component';

@NgModule({
  imports: [BrowserModule],
  declarations: [AppComponent, CurrentTimeComponent],
  providers: [TimeService],
  bootstrap: [AppComponent]
})
export class AppModule { }
