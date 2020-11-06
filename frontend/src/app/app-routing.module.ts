import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { StatsComponent } from './stats/stats.component';
import { ConfigFormComponent } from './keyboard/config-form/config-form.component';
import { KeyboardComponent } from './keyboard/keyboard.component';
import { ConfigListComponent } from './keyboard/config-list/config-list.component';
import { RunningConfigComponent } from './keyboard/running-config/running-config.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'stats' },
  { path: 'stats', component: StatsComponent },
  { path: 'keyboard', component: KeyboardComponent, children: [
      { path: '', redirectTo: 'list', pathMatch: 'full'},
      { path: 'list', component: ConfigListComponent },
      { path: 'edit/config/:id', component: ConfigFormComponent },
      { path: 'running', component: RunningConfigComponent }
    ]
  },
  { path: '**', pathMatch: 'full', redirectTo: 'stats' }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {useHash: true})
  ],
  exports: [RouterModule]
})

export class AppRoutingModule { }
