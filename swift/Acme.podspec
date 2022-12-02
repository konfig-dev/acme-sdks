Pod::Spec.new do |s|
  s.name = 'Acme'
  s.ios.deployment_target = '9.0'
  s.osx.deployment_target = '10.11'
  s.tvos.deployment_target = '9.0'
  s.watchos.deployment_target = '3.0'
  s.version = '1.0.0'
  s.source = { :git => 'git@github.com:OpenAPITools/openapi-generator.git', :tag => 'v1.0.0' }
  s.authors = 'acme.com'
  s.license = 'Proprietary'
  s.homepage = 'https://github.com/OpenAPITools/openapi-generator'
  s.summary = 'Acme Swift SDK'
  s.source_files = 'Acme/**/*.swift'
  s.dependency 'AnyCodable-FlightSchool', '~> 0.6.1'
end
