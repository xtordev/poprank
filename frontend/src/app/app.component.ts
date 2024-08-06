import { Component, Injectable, OnInit } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';
import { HeaderComponent } from "./components/header/header.component";
import { HttpClient } from '@angular/common/http';
import { Student } from './interface/student.interface';
import { MatchService } from './services/match.service';
@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, HeaderComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {
  constructor(private http:HttpClient, private readonly matchService:MatchService, private router:Router) {
  }
  title = 'frontend';
  match:{
    student1:Student,
    student2:Student
  } | undefined
  loading:boolean=false
  ngOnInit(): void {
    this.matchService.getMatch().subscribe((data) => {this.match = data});
  }
  getMatchData(): void {
    this.matchService.getMatch().subscribe((data) => {
      this.match = data;
      this.loading = false;
    });

  }
  matchWin(student1:number, student2:number, matchWinner:number){
    if (this.loading) {return}
    this.loading = true;
    this.matchService.matchWin(student1,student2,matchWinner).subscribe((data) => {this.match = data;
      this.getMatchData();
    });

  }
}

