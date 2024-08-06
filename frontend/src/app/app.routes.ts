import { Routes } from '@angular/router';
import { LeaderboardComponent } from './pages/leaderboard/leaderboard.component';
import { HomeComponent } from './pages/home/home.component';

export const routes: Routes = [
  {path: "", component:HomeComponent},
  {path:"leaderboard",component:LeaderboardComponent}
];
