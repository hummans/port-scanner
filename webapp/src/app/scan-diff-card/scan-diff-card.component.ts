import { Component, OnInit, Input } from '@angular/core';
import { ScanResult } from '../scan';


@Component({
  selector: 'app-scan-diff-card',
  templateUrl: './scan-diff-card.component.html',
  styleUrls: ['./scan-diff-card.component.css']
})
export class ScanDiffCardComponent implements OnInit {
  @Input() scanResult?: ScanResult;

  constructor() { }

  ngOnInit(): void { }
}
