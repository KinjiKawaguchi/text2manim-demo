import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class GenerationService {
	private apiUrl = environment.apiEndpoint + '/v1/generations';

	constructor(private http: HttpClient) { }

	sendGenerationRequest(prompt: string, email: string): Observable<{ request_id: string }> {
		const body = { prompt, email };
		return this.http.post<{ request_id: string }>(this.apiUrl, body);
	}

	getGenerationStatus(requestId: string): Observable<{ status: string }> {
		return this.http.get<{ status: string }>(`${this.apiUrl}/${requestId}`);
	}
}
