import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Student } from '../interface/student.interface';

@Injectable({
  providedIn: 'root'
})
export class MatchService {

  constructor(private http:HttpClient) { }

  getMatch() {
    return this.http.get<{
      student1:Student,
      student2:Student
    }>(`${import.meta.env.NG_APP_BACKEND_URL}/matches`,{responseType:'json'})
  }

  matchWin(student1:number,student2:number,matchWinner:number){
    return this.http.post<{student1:Student,student2:Student}>(`${import.meta.env.NG_APP_BACKEND_URL}/matches/win`,{
      StudentId1:student1,
      StudentId2:student2,
      MatchWinner:matchWinner
    },{responseType:"json"})
  }
}
