import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { HttpClient, HttpHeaders, HttpParams, HttpErrorResponse } from '@angular/common/http';
import { environment } from '../environments/environment';
import { Scan, ScanResult } from './scan';

@Injectable({
  providedIn: 'root'
})
export class ScannerService {

  private apiBase;

  constructor(private http: HttpClient) { 
    this.apiBase = environment.apiBase;
  }

  private handleError(error: HttpErrorResponse) {
    return throwError(error.error);
  }

  ListScans(host: string): Observable<Scan[]> {
    const params = new HttpParams().set('host', host);
    const url = `${this.apiBase}/scans`;

    return this.http.get<Scan[]>(url, { params: params })
      .pipe(catchError(this.handleError));
  }

  Scan(host: string): Observable<ScanResult> {
    const params = new HttpParams().set('host', host);
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const url = `${this.apiBase}/scans`;

    return this.http.post<ScanResult>(url, null, { params: params, headers: headers })
      .pipe(catchError(this.handleError));  
    }
}