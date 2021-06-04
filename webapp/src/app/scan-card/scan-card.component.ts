import { Component, OnInit, Input } from '@angular/core';
import { Scan } from '../scan';

@Component({
  selector: 'app-scan-card',
  templateUrl: './scan-card.component.html',
  styleUrls: ['./scan-card.component.css']
})
export class ScanCardComponent implements OnInit {
  @Input() scan?: Scan;

  constructor() { }
  ngOnInit(): void { }
}