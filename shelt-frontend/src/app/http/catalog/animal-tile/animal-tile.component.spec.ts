import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnimalTileComponent } from './animal-tile.component';

describe('AnimalTileComponent', () => {
  let component: AnimalTileComponent;
  let fixture: ComponentFixture<AnimalTileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnimalTileComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AnimalTileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
