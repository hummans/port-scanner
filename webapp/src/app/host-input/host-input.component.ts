import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { Scan, ScanResult } from '../scan';
import { ScannerService } from '../scanner.service';

@Component({
  selector: 'app-host-input',
  templateUrl: './host-input.component.html',
  styleUrls: ['./host-input.component.css']
})
export class HostInputComponent implements OnInit {

  host: string = '';
  error: string = '';
  loading: boolean = false;

  viewHistory: boolean = false;
  scanHistory: Scan[] = [];
  scanResult?: ScanResult;

  constructor(private scanner: ScannerService, private router: Router) { }

  ngOnInit(): void { }

  onInput(val: string): void {
    // REVIEW: This would be a good place to plug in realtime input validation.
    this.host = val;
  }

  scan(): void {
    this.loading = true;
    this.viewHistory = false;
    this.scanHistory = [];
    this.error = '';

    this.scanner.Scan(this.host)
      .subscribe(result => {
        this.scanResult = result;
        this.loading = false;
      }, (error) => {
        this.error = error.error;
        this.loading = false;
      });
  }

  history(): void {
    this.loading = true;
    this.viewHistory = true;
    this.scanResult = undefined;
    this.error = '';

    this.scanner.ListScans(this.host)
      .subscribe(result => {
        this.scanHistory = result;
        this.loading = false;
      }, (error) => {
        this.error = error.error;
        this.loading = false;
      });
  }
}
