import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { ScanCardComponent } from './scan-card/scan-card.component';
import { ScanDiffCardComponent } from './scan-diff-card/scan-diff-card.component';
import { HostInputComponent } from './host-input/host-input.component';
import { AboutComponent } from './about/about.component';

@NgModule({
  declarations: [
    AppComponent,
    HostInputComponent,
    ScanCardComponent,
    ScanDiffCardComponent,
    AboutComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
