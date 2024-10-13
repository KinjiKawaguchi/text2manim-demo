import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

interface GenerationData {
	ID: number;
	RequestID: string;
	Email: string;
	Prompt: string;
	Status: string;
	VideoURL: string;
	ScriptURL: string;
	ErrorMessage: string;
	CreatedAt: string;
	UpdatedAt: string;
}

interface GenerationResponse {
	generation_status: GenerationData;
}

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


	getGenerationStatus(requestId: string): Observable<GenerationData> {
		return this.http.get<GenerationResponse>(`${this.apiUrl}/${requestId}`).pipe(
			map(response => response.generation_status)
		);
	}
}
