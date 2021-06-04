export interface Scan {
    id: number;
    created_at: string;
    host: string;
    ports: Port[];
}

export interface Port {
    port: number;
    protocol: string;
    state: string;
    name: string;
}

export interface Diff {
    from: string;
    to: string;
}

export interface ScanResult {
    diff: Diff[];
    scan: Scan;
}