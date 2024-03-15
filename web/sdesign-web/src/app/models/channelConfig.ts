import { InputConfig } from "./inputConfig";
import { EQConfig } from "./eqConfig";
import { CompressorConfig } from "./compressorConfig";
import { GateConfig } from "./gateConfig";

export interface ChannelConfig {
    name: string;
    id: number;
    color: string;
    input_cfg: InputConfig;
    eq_cfg: EQConfig;
    compressor_cfg: CompressorConfig;
    gate_cfg: GateConfig;
}
