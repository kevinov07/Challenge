
export interface Response {
took: number;
timed_out: boolean;
_shards: Shards;
hits: Hits;
}

export interface Shards {
total: number;
successful: number;
skipped: number;
failed: number;
}

export interface Hit {
_index: string;
_type: string;
_id: string;
_score: number;
"@timestamp": string;
_source: Email;
}

export interface Hits {
total: Total;
max_score: number;
hits: Hit[];
}

export interface Email {
"@timestamp": string;
body: string;
c_folder: string;
cc: string;
content_transfer_encoding: string;
content_type: string;
date: string;
date_subemail: string;
from: string;
message_id: string;
mime_version: string;
sent: string;
subject: string;
to: string;
x_bcc: string;
x_cc: string;
x_file_name: string;
x_from: string;
x_origin: string;
x_to: string;
}

export interface Total {
    value: number;

} 