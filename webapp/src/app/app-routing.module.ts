import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HostInputComponent } from './host-input/host-input.component';
import { AboutComponent } from './about/about.component';

const routes: Routes = [
  { path: '', component: HostInputComponent },
  { path: 'about', component: AboutComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
