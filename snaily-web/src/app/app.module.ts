import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {AppComponent} from './app.component';
import {FormsModule} from '@angular/forms';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RegisterService} from './services/register.service';
import {HttpClientModule} from '@angular/common/http';

// Libs
import {HttpModule} from '@angular/http';

// Routing
import {AppRoutingModule, routedComponents} from './app-routing.module';
import {WelcomeComponent} from './components/welcome/welcome.component';
import {HomeComponent} from './components/home/home.component';
import {NewComponent} from './components/new/new.component';
import {HelpComponent} from './components/help/help.component';
import {ViewComponent} from './components/view/view.component';
import {HttpClient} from '@angular/common/http';
import {LoginService} from './services/login.service';
import {ItemService} from './services/item.service';

// import './rxjs-extensions';

@NgModule({
  declarations: [
    AppComponent,
    routedComponents,
    WelcomeComponent,
    HomeComponent,
    NewComponent,
    ViewComponent,
    HelpComponent
  ],
  imports: [
    NgbModule.forRoot(),
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpModule,
    HttpClientModule,

  ],
  providers: [RegisterService, HttpClient, LoginService, ItemService],
  bootstrap: [AppComponent]
})
export class AppModule {
}
