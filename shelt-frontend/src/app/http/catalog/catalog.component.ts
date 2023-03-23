import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-catalog',
  templateUrl: './catalog.component.html',
  styleUrls: ['./catalog.component.scss']
})
export class CatalogComponent implements OnInit {

  animals: any;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.http.get("http://localhost:8080/animals?fields=name,gender,shoulderheight,Description").subscribe( (data) => {
      this.animals = data;
      console.log(this.animals[0])
    })
  }

}
