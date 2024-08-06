import { Component, OnInit } from '@angular/core';
import { HeaderComponent } from "../../components/header/header.component";
import { Student } from '../../interface/student.interface';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-leaderboard',
  standalone: true,
  imports: [HeaderComponent],
  templateUrl: './leaderboard.component.html',
  styleUrl: './leaderboard.component.css'
})
export class LeaderboardComponent implements OnInit{
  constructor(private http:HttpClient) { }
  leaderboard:Student[]|undefined

  ngOnInit(): void {
      this.http.get<Student[]>(`${import.meta.env.NG_APP_BACKEND_URL}/leaderboard`)
     .subscribe(data => this.leaderboard = data);
  }
}
