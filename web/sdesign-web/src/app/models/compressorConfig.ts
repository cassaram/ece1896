export interface CompressorConfig {
  bypass: boolean;
  pre_gain: number;
  post_gain: number;
  threshold: number;
  ratio: number;
  attack_time: number;
  release_time: number;
}
