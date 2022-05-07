

export interface Menu {
    name: string;
    path: string;
    icon: string;
    components: string;
    hidden?: boolean;
    parent_id: number;
    redirect?: string;
    label?: string;
}
