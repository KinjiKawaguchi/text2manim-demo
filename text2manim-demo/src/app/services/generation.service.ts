import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

const apiEndpoint = 'https://api.text2manim-demo.kawakin.tech';

interface GenerationData {
	ID: number;
	RequestID: string;
	Email: string;
	Prompt: string;
	Status: string;
	VideoUrl: string;
	ScriptUrl: string;
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
	private apiUrl = apiEndpoint + '/v1/generations';

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
