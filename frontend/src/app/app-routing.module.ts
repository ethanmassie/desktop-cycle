import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { StatsComponent } from './stats/stats.component';
import { KeyboardConfigComponent } from './keyboard-config/keyboard-config.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'stats' },
  { path: 'stats', component: StatsComponent },
  { path: 'keyboard', component: KeyboardConfigComponent },
  { path: '**', pathMatch: 'full', redirectTo: 'stats' }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {useHash: true})
  ],
  exports: [RouterModule]
})

export class AppRoutingModule { }
