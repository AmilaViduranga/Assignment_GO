import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { UtilsModule } from './utils/utils.module';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { AppRouting } from './app.route';
import { AppService } from './app.service';

import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    AppRouting,
    UtilsModule,
    HttpClientModule,
  ],
  providers: [
    AppService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
