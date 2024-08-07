import { Component, OnInit } from '@angular/core';
import { HeaderComponent } from "../../components/header/header.component";
import { Student } from '../../interface/student.interface';
import { HttpClient } from '@angular/common/http';
import { LeaderboardService } from '../../services/leaderboard.service';

@Component({
  selector: 'app-leaderboard',
  standalone: true,
  imports: [HeaderComponent],
  templateUrl: './leaderboard.component.html',
  styleUrl: './leaderboard.component.css'
})
export class LeaderboardComponent implements OnInit{
  constructor(private leaderboardService:LeaderboardService) { }
  leaderboard:Student[]|undefined

  ngOnInit(): void {
     this.leaderboardService.getLeaderboard().subscribe((data) => {this.leaderboard = data});
  }
}
