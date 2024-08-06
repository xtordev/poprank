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
export class AppComponent{
  constructor(private http:HttpClient, private readonly matchService:MatchService, private router:Router) {
  }
  title = 'frontend';

}

