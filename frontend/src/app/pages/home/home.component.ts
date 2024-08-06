import { Component, OnInit } from '@angular/core';
import { MatchService } from '../../services/match.service';
import { Student } from '../../interface/student.interface';
import { HeaderComponent } from '../../components/header/header.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [HeaderComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent  {
  constructor(private matchService:MatchService){}
  match:{
    student1:Student,
    student2:Student
  } | undefined
  loading:boolean=false
  OnInit(): void {
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
