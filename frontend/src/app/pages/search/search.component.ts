import { Component } from '@angular/core';
import { HeaderComponent } from "../../components/header/header.component";
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Student } from '../../interface/student.interface';
import { HttpClient } from '@angular/common/http';

interface StudentInfo{
  percentile:number,
  ranking:number,
  student:Student
}

@Component({
  selector: 'app-search',
  standalone: true,
  imports: [HeaderComponent, ReactiveFormsModule],
  templateUrl: './search.component.html',
  styleUrl: './search.component.css'
})
export class SearchComponent {
  searchForm = new FormGroup({
    Code: new FormControl("")
  })
  studentInfo:StudentInfo|undefined
  loading:boolean=false
  error:boolean=false
  errMsg:string=""
  constructor(private http:HttpClient) { }

  getStudent( ) {
    this.loading = true;
    this.error = false;
    this.errMsg = "";
    const code = this.searchForm.get('Code')?.value;
    this.http.post<StudentInfo>(`${import.meta.env.NG_APP_BACKEND_URL}/leaderboard/standing`,{Code:code},{responseType:"json"})
    .subscribe(
      (data) => this.studentInfo = data,
      (error) =>{this.error = true;
        this.studentInfo = undefined
        if (error.status == 400) {
          this.errMsg = "Invalid Student Code"
        }
        else if (error.status == 404) {
          this.errMsg = error.error.error;}
        else {
          this.errMsg = "Server is not responding"
          }
        }
    )
    this.loading = false;
  }
}
