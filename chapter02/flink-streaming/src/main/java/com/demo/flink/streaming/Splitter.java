package com.demo.flink.streaming;

import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.util.Collector;
import org.apache.flink.util.LinkedOptionalMap;
import org.apache.log4j.Logger;

public class Splitter implements FlatMapFunction<String, Tuple2<String, Double>> {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1L;
	private static Logger logger = Logger.getLogger(Splitter.class);
	@Override
	public void flatMap(String value, Collector<Tuple2<String, Double>> out) throws Exception {
//		logger.info(String.format("%s, %s, %s",value,!"".equals(value),value.contains(",")));
		if (!"".equals(value) && value.contains(",")) {
			String parts[] = value.split(",");
			out.collect(new Tuple2<String, Double>(parts[2], Double.parseDouble(parts[1])));
//			logger.info("============================");
//            logger.info(String.format("%s, %s", parts[2], parts[1]));
		}
	}
}
