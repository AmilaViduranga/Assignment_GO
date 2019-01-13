import { RouterModule, Routes } from "@angular/router";
import { HomeComponent } from "./home/home.component";
import { LoginComponent } from './login/login.component';

const appRoutes:Routes = [
    { path: 'login', component: LoginComponent },
    { path: 'home', component: HomeComponent },
    { path: '', redirectTo: 'home', pathMatch: 'full' }
]

export const AppRouting = RouterModule.forRoot(appRoutes, {
    useHash: true
});