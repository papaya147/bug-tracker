type Bug = {
  id: string;
  name: string;
  description: string;
  status: string;
  priority: string;
  assigned_to: string;
  assigned_by_profile: string;
  assigned_by_team: string;
  completed: boolean;
  createdat: number;
  updated_at: number;
  closed_by?: string;
  remarks?: string;
  closed_at?: number;
};

export default Bug;
