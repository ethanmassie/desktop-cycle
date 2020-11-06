import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { APP_BASE_HREF } from '@angular/common';
import { StatsComponent } from './stats/stats.component';
import { StoreModule } from '@ngrx/store';
import { reducerMap } from './state';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ConfigFormComponent } from './keyboard/config-form/config-form.component';
import { KeyboardComponent } from './keyboard/keyboard.component';
import { ConfigListComponent } from './keyboard/config-list/config-list.component';
import { RunningConfigComponent } from './keyboard/running-config/running-config.component';
import { MatSelectModule, MatTableModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { CoreModule } from './core/core.module';

@NgModule({
  declarations: [
    AppComponent,
    StatsComponent,
    ConfigFormComponent,
    KeyboardComponent,
    ConfigListComponent,
    RunningConfigComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    CoreModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    MatTableModule,
    MatSelectModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    StoreModule.forRoot(reducerMap)
  ],
  providers: [{provide: APP_BASE_HREF, useValue : '/' }],
  bootstrap: [AppComponent]
})
export class AppModule { }
