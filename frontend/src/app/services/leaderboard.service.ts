import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Student } from '../interface/student.interface';

@Injectable({
  providedIn: 'root'
})
export class LeaderboardService {

  constructor(private http:HttpClient) { }

  getLeaderboard() {
    return  this.http.get<Student[]>(`${import.meta.env.NG_APP_BACKEND_URL}/leaderboard`)
  }
}
