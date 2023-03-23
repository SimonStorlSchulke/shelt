import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AnimalTileComponent } from './http/catalog/animal-tile/animal-tile.component';
import { CatalogComponent } from './http/catalog/catalog.component';

@NgModule({
  declarations: [
    AppComponent,
    CatalogComponent,
    AnimalTileComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
